-- name: CreateProductVersion :one
INSERT INTO product_versions (
    id, product_id, branch_id, name, 
    price_adjustment, attributes, stock_quantity, reorder_point
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetProductVersionsByProductID :many
SELECT * FROM product_versions WHERE product_id = $1;


-- name: UpdateProductVersion :one
UPDATE product_versions 
SET 
    name = $2, 
    price_adjustment = $3,
    attributes = $4,
    stock_quantity = $5,
    reorder_point = $6
WHERE id = $1 
RETURNING *;