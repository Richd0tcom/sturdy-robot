-- name: CreateInvoiceItem :one
INSERT INTO invoice_items (
    id, invoice_id, version_id, quantity, 
    unit_price, subtotal
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateMultipleInvoiceItems :copyfrom
INSERT INTO invoice_items (
    id,
    invoice_id, version_id, quantity, 
    unit_price, subtotal
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetInvoiceItemsByInvoiceID :many
SELECT * FROM invoice_items it
JOIN product_versions pv ON pv.id = it.version_id
WHERE it.invoice_id = $1;

-- name: UpdateInvoiceItem :one
UPDATE invoice_items 
SET 
    version_id = $2, 
    quantity = $3, 
    unit_price = $4, 
    subtotal = $5
WHERE id = $1 
RETURNING *;

-- name: DeleteItemsByInvoiceId :exec
DELETE FROM invoice_items
WHERE invoice_id = $1;