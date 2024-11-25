-- Insert payment info
-- name: CreatePaymentInfo :one
INSERT INTO payment_info (id, user_id, account_no, routing_no, account_name, bank_name)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5) RETURNING *;

-- name: GetPaymentInfoByUserID :one
SELECT * FROM payment_info WHERE user_id = $1 LIMIT 1;

-- name: UpdatePaymentInfo :exec
UPDATE payment_info 
SET account_no = $2, routing_no = $3, account_name = $4, bank_name = $5
WHERE user_id = $1;