-- name: CreateCurrency :one
INSERT INTO currency (id, name, symbol, code) 
VALUES (uuid_generate_v4(), $1, $2, $3) 
RETURNING *;

-- name: GetCurrencyByID :one
SELECT * FROM currency 
WHERE id = $1 LIMIT 1;