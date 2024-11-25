-- name: CreatePayment :one
INSERT INTO payments (
    id, invoice_id, payment_method, payment_amount, 
    payment_ref, payment_date, metadata, created_by
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- Get payment by ID
-- name: GetPaymentByID :one
SELECT * FROM payments WHERE id = $1;

-- Get payments by invoice ID
-- name: GetPaymentsByInvoiceID :many
SELECT * FROM payments WHERE invoice_id = $1;

-- Update payment
-- name: UpdatePayment :one
UPDATE payments 
SET 
    payment_method = $2, payment_amount = $3, 
    payment_ref = $4, payment_date = $5, metadata = $6
WHERE id = $1 RETURNING *;