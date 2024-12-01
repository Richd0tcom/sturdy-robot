-- name: CreateInvoice :one
INSERT INTO invoices (
    id, customer_id, invoice_number, subtotal, 
    discount, total, status, created_by, 
    currency_id, due_date, reminders, payment_info
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, 
    $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: GetInvoiceByID :one
SELECT * FROM invoices WHERE id = $1 LIMIT 1;

-- name: GetInvoicesCreatedByUser :many
SELECT * FROM invoices WHERE created_by = $1;

-- name: UpdateInvoice :one
UPDATE invoices 
SET 
    customer_id = $2, 
    subtotal = $3, 
    discount = $4, 
    total = $5, 
    status = $6, 
    reminders = $7,
     currency_id= $8,
     metadata= $9,
     due_date= $10,
     payment_info= $11
WHERE id = $1 
RETURNING *;

-- name: GetTotalsByStatuses :one
SELECT
    SUM(CASE WHEN status = 'paid' THEN total ELSE 0 END) AS paid_total,
    SUM(CASE WHEN status = 'unpaid' THEN total ELSE 0 END) AS unpaid_total,
    SUM(CASE WHEN status = 'overdue' THEN total ELSE 0 END) AS overdue_total,
    SUM(CASE WHEN status = 'draft' THEN total ELSE 0 END) AS draft_total
FROM
    invoices
WHERE
    created_by = $1;


-- name: UpdateInvoicePayment :one
UPDATE invoices 
SET 
    amount_paid = $2, 
    status = $3
WHERE id = $1 
RETURNING *;


