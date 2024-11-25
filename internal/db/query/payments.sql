-- name: CreatePayment
INSERT INTO payments (
    id, invoice_id, payment_method, payment_amount, 
    payment_ref, payment_date, metadata, created_by
) VALUES (
    nextval('payments_id_seq'), $1, $2, $3, $4, $5, $6, $7
);

-- Get payment by ID
-- name: GetPaymentByID
SELECT * FROM payments WHERE id = $1;

-- Get payments by invoice ID
-- name: GetPaymentsByInvoiceID
SELECT * FROM payments WHERE invoice_id = $1;

-- Update payment
-- name: UpdatePayment
UPDATE payments 
SET 
    payment_method = $2, payment_amount = $3, 
    payment_ref = $4, payment_date = $5, metadata = $6
WHERE id = $1;