-- name: CreateInvoice :one
INSERT INTO invoices (
    id, customer_id, invoice_number, subtotal, 
    discount, total, status, created_by, 
    currency_id, due_date, reminders
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, 
    $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetInvoiceByID :one
SELECT * FROM invoices WHERE id = $1 LIMIT 1;

-- name: GetInvoicesCreatedByUser :many
SELECT * FROM invoices WHERE created_by = $1;

-- name: UpdateInvoice :one
UPDATE invoices 
SET 
    invoice_number = $2, 
    subtotal = $3, 
    discount = $4, 
    total = $5, 
    status = $6, 
    amount_paid = $7,
    reminders = $8
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
