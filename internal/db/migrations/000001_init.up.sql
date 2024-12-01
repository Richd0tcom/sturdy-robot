-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE product_type AS ENUM ('physical', 'service'); -- can be created oon the backend


-- Organizations table
CREATE TABLE organizations (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


-- Branches table
CREATE TABLE branches (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    is_default BOOLEAN DEFAULT true,
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--users table 
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT,
    branch_id UUID NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--payment_info table 
CREATE TABLE payment_info (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    account_no VARCHAR(255) NOT NULL,
    routing_no VARCHAR(255),
    account_name VARCHAR(255) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);



-- Categories table with self-referential relationship
CREATE TABLE categories (
    id UUID PRIMARY KEY,
    parent_id UUID REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    branch_id UUID NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT categories_prevent_self_loop CHECK (id != parent_id)
);

-- Products table
CREATE TABLE products (
    id UUID PRIMARY KEY,
    category_id UUID NOT NULL REFERENCES categories(id),
    branch_id UUID NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    product_type VARCHAR(50) NOT NULL DEFAULT 'physical',
    service_pricing_model VARCHAR(50), --hourly, per-project, tiered
    default_unit VARCHAR(50),-- hours, projects, sessions
    is_billable BOOLEAN DEFAULT true,
    sku VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    base_price DECIMAL(15,2),
    custom_fields JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Product versions table
CREATE TABLE product_versions (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    branch_id UUID NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    price_adjustment DECIMAL(15,2) DEFAULT 0,
    attributes JSONB,
    stock_quantity INTEGER DEFAULT 0,
    reorder_point INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Inventory table
CREATE TABLE inventory (
    id UUID PRIMARY KEY,
    version_id UUID NOT NULL REFERENCES product_versions(id),
    branch_id UUID NOT NULL REFERENCES branches(id),
    quantity INTEGER,
    unit_cost DECIMAL(15,2),
    last_counted TIMESTAMP WITH TIME ZONE
    -- CONSTRAINT unique_version_branch UNIQUE(version_id, branch_id),
--     CONSTRAINT valid_quantity CHECK (
--     -- Only require quantity for physical products
--     (quantity IS NOT NULL AND quantity >= 0) = EXISTS (
--         SELECT 1 FROM products p 
--         JOIN product_versions pv ON p.id = pv.product_id 
--         WHERE pv.id = inventory.version_id 
--         AND p.product_type = 'physical'
--     )
-- )
);

-- Customers table
CREATE TABLE customers (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50),
    billing_address JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    branch_id UUID NOT NULL REFERENCES branches(id)
);

-- Currency table
CREATE TABLE currency (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) NOT NULL,
    symbol VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);



-- Invoices table
CREATE TABLE invoices (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES customers(id),
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    subtotal DECIMAL(15,2) NOT NULL DEFAULT 0,
    discount DECIMAL(15,2) DEFAULT 0,
    total DECIMAL(15,2) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'draft',
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    currency_id UUID REFERENCES currency(id),

    due_date TIMESTAMP WITH TIME ZONE,
    reminders JSONB,
    metadata JSONB,

    -- for patial payment
    amount_paid DECIMAL(15,2) NOT NULL DEFAULT 0,
    balance_due DECIMAL(15,2) GENERATED ALWAYS AS (total - amount_paid) STORED,
    CONSTRAINT valid_total CHECK (total = subtotal - discount),

    -- sender payment info 
    payment_info UUID REFERENCES payment_info(id)
);

-- Payment transactions table
CREATE TABLE payments (
    id UUID PRIMARY KEY,
    invoice_id UUID REFERENCES invoices(id),          -- Links to specific invoice
    payment_method VARCHAR(255), 
    payment_amount DECIMAL(10,2) NOT NULL,
    payment_ref VARCHAR(255),
    payment_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_by UUID NOT NULL -- user who recorded the payment
);

-- Invoice items table
CREATE TABLE invoice_items (
    id UUID PRIMARY KEY,
    invoice_id UUID NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
    version_id UUID NOT NULL REFERENCES product_versions(id),
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(15,2) NOT NULL,
    subtotal DECIMAL(15,2) NOT NULL,
    metadata JSONB,
    CONSTRAINT positive_quantity CHECK (quantity > 0),
    CONSTRAINT positive_price CHECK (unit_price >= 0)
    -- CONSTRAINT valid_subtotal CHECK (subtotal = quantity * unit_price)
);

-- Activity logs table
CREATE TABLE activity_logs (
    id UUID PRIMARY KEY,
    entity_type VARCHAR(50) NOT NULL,
    entity_id UUID NOT NULL,
    action VARCHAR(50) NOT NULL,
    changes JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    user_id UUID 
);

-- Indexes for better query performance
CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_versions_product ON product_versions(product_id);
CREATE INDEX idx_inventory_version ON inventory(version_id);
CREATE INDEX idx_inventory_branch ON inventory(branch_id);
CREATE INDEX idx_invoices_customer ON invoices(customer_id);
CREATE INDEX idx_invoice_items_invoice ON invoice_items(invoice_id);
CREATE INDEX idx_invoice_items_version ON invoice_items(version_id);
CREATE INDEX idx_activity_logs_entity ON activity_logs(entity_type, entity_id);
CREATE INDEX idx_categories_parent ON categories(parent_id);
CREATE INDEX idx_payments_invoice ON payments(invoice_id);

-- Auto-update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER 
LANGUAGE plpgsql
AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;

CREATE TRIGGER update_products_updated_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
