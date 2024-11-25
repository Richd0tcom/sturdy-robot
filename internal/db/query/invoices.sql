-- name: CreateInvoice :one
INSERT INTO invoices (
    id, customer_id, invoice_number, subtotal, 
    discount, total, status, created_by, 
    currency_id, due_date
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, 
    $6, $7, $8, $9
) RETURNING *;

-- name: GetInvoiceByID :one
SELECT * FROM invoices WHERE id = $1 LIMIT 1;

-- name: GetInvoicesCreatedByUser
SELECT * FROM invoices WHERE created_by = $1;

-- name: UpdateInvoice
UPDATE invoices 
SET 
    invoice_number = $2, 
    subtotal = $3, 
    discount = $4, 
    total = $5, 
    status = $6, 
    amount_paid = $7
WHERE id = $1 
RETURNING *;
