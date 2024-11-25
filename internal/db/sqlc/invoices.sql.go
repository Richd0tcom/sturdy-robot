// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: invoices.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO invoices (
    id, customer_id, invoice_number, subtotal, 
    discount, total, status, created_by, 
    currency_id, due_date
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, 
    $6, $7, $8, $9
) RETURNING id, customer_id, invoice_number, subtotal, discount, total, status, created_by, created_at, currency_id, due_date, reminders, metadata, amount_paid, balance_due
`

type CreateInvoiceParams struct {
	CustomerID    uuid.NullUUID `json:"customer_id"`
	InvoiceNumber string        `json:"invoice_number"`
	Subtotal      string        `json:"subtotal"`
	Discount      string        `json:"discount"`
	Total         string        `json:"total"`
	Status        string        `json:"status"`
	CreatedBy     uuid.NullUUID `json:"created_by"`
	CurrencyID    uuid.NullUUID `json:"currency_id"`
	DueDate       sql.NullTime  `json:"due_date"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice,
		arg.CustomerID,
		arg.InvoiceNumber,
		arg.Subtotal,
		arg.Discount,
		arg.Total,
		arg.Status,
		arg.CreatedBy,
		arg.CurrencyID,
		arg.DueDate,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.InvoiceNumber,
		&i.Subtotal,
		&i.Discount,
		&i.Total,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.CurrencyID,
		&i.DueDate,
		&i.Reminders,
		&i.Metadata,
		&i.AmountPaid,
		&i.BalanceDue,
	)
	return i, err
}

const getInvoiceByID = `-- name: GetInvoiceByID :one
SELECT id, customer_id, invoice_number, subtotal, discount, total, status, created_by, created_at, currency_id, due_date, reminders, metadata, amount_paid, balance_due FROM invoices WHERE id = $1 LIMIT 1
`

func (q *Queries) GetInvoiceByID(ctx context.Context, id uuid.UUID) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, getInvoiceByID, id)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.InvoiceNumber,
		&i.Subtotal,
		&i.Discount,
		&i.Total,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.CurrencyID,
		&i.DueDate,
		&i.Reminders,
		&i.Metadata,
		&i.AmountPaid,
		&i.BalanceDue,
	)
	return i, err
}

const getInvoicesCreatedByUser = `-- name: GetInvoicesCreatedByUser :many
SELECT id, customer_id, invoice_number, subtotal, discount, total, status, created_by, created_at, currency_id, due_date, reminders, metadata, amount_paid, balance_due FROM invoices WHERE created_by = $1
`

func (q *Queries) GetInvoicesCreatedByUser(ctx context.Context, createdBy uuid.NullUUID) ([]Invoice, error) {
	rows, err := q.db.QueryContext(ctx, getInvoicesCreatedByUser, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Invoice{}
	for rows.Next() {
		var i Invoice
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.InvoiceNumber,
			&i.Subtotal,
			&i.Discount,
			&i.Total,
			&i.Status,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.CurrencyID,
			&i.DueDate,
			&i.Reminders,
			&i.Metadata,
			&i.AmountPaid,
			&i.BalanceDue,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInvoice = `-- name: UpdateInvoice :one
UPDATE invoices 
SET 
    invoice_number = $2, 
    subtotal = $3, 
    discount = $4, 
    total = $5, 
    status = $6, 
    amount_paid = $7
WHERE id = $1 
RETURNING id, customer_id, invoice_number, subtotal, discount, total, status, created_by, created_at, currency_id, due_date, reminders, metadata, amount_paid, balance_due
`

type UpdateInvoiceParams struct {
	ID            uuid.UUID `json:"id"`
	InvoiceNumber string    `json:"invoice_number"`
	Subtotal      string    `json:"subtotal"`
	Discount      string    `json:"discount"`
	Total         string    `json:"total"`
	Status        string    `json:"status"`
	AmountPaid    string    `json:"amount_paid"`
}

func (q *Queries) UpdateInvoice(ctx context.Context, arg UpdateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, updateInvoice,
		arg.ID,
		arg.InvoiceNumber,
		arg.Subtotal,
		arg.Discount,
		arg.Total,
		arg.Status,
		arg.AmountPaid,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.InvoiceNumber,
		&i.Subtotal,
		&i.Discount,
		&i.Total,
		&i.Status,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.CurrencyID,
		&i.DueDate,
		&i.Reminders,
		&i.Metadata,
		&i.AmountPaid,
		&i.BalanceDue,
	)
	return i, err
}
