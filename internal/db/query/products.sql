-- name: CreateProduct :one
INSERT INTO products (
    id, category_id, branch_id, name, product_type, 
    service_pricing_model, default_unit, is_billable, 
    sku, description, base_price
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetProductsByID :one
SELECT * FROM products 
WHERE id = $1 LIMIT 1;


-- name: GetProductsByBranchID :one
SELECT * FROM products 
WHERE branch_id = $1 LIMIT 1;

-- name: UpdateProduct :one
UPDATE products 
SET 
    category_id = $2, 
    name = $3, 
    product_type = $4, 
    service_pricing_model = $5,
    default_unit = $6,
    is_billable = $7,
    sku = $8,
    description = $9,
    base_price = $10
WHERE id = $1 
RETURNING *;