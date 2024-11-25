DROP TRIGGER IF EXISTS update_products_updated_at ON products;
DROP FUNCTION IF EXISTS update_updated_at_column();

DROP INDEX IF EXISTS idx_payments_invoice;
DROP INDEX IF EXISTS idx_categories_parent;
DROP INDEX IF EXISTS idx_activity_logs_entity;
DROP INDEX IF EXISTS idx_invoice_items_version;
DROP INDEX IF EXISTS idx_invoice_items_invoice;
DROP INDEX IF EXISTS idx_invoices_customer;
DROP INDEX IF EXISTS idx_inventory_branch;
DROP INDEX IF EXISTS idx_inventory_version;
DROP INDEX IF EXISTS idx_versions_product;
DROP INDEX IF EXISTS idx_products_category;

DROP TABLE IF EXISTS activity_logs;
DROP TABLE IF EXISTS invoice_items;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS invoices;
DROP TABLE IF EXISTS currency;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS product_versions;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS payment_info;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS branches;
DROP TABLE IF EXISTS organizations;

DROP TYPE IF EXISTS product_type;

DROP EXTENSION IF EXISTS "uuid-ossp";