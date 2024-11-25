-- name: CreateProductVersion :one
INSERT INTO product_versions (
    id, product_id, branch_id, sku, name, 
    price_adjustment, attributes, stock_quantity, reorder_point
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetProductVersionsByProductID :many
SELECT * FROM product_versions WHERE product_id = $1;


-- name: UpdateProductVersion :one
UPDATE product_versions 
SET 
    sku = $2, 
    name = $3, 
    price_adjustment = $4,
    attributes = $5,
    stock_quantity = $6,
    reorder_point = $7
WHERE id = $1 
RETURNING *;