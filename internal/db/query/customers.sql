-- name: CreateCustomer :one
INSERT INTO customers (id, name, email, phone, billing_address) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4) 
RETURNING *;

-- name: GetCustomerByEmail :one
SELECT * FROM customers 
WHERE email = $1 LIMIT 1;

-- name: GetCustomerById :one
SELECT * FROM customers 
WHERE id = $1 LIMIT 1;

-- name: DeleteCustomerByID :exec
DELETE FROM customers
WHERE id = $1;