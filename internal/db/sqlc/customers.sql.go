// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: customers.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (id, name, email, phone, billing_address) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4) 
RETURNING id, name, email, phone, billing_address, created_at
`

type CreateCustomerParams struct {
	Name           string      `json:"name"`
	Email          pgtype.Text `json:"email"`
	Phone          pgtype.Text `json:"phone"`
	BillingAddress []byte      `json:"billing_address"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, createCustomer,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.BillingAddress,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.BillingAddress,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCustomerByID = `-- name: DeleteCustomerByID :exec
DELETE FROM customers
WHERE id = $1
`

func (q *Queries) DeleteCustomerByID(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCustomerByID, id)
	return err
}

const getCustomerByEmail = `-- name: GetCustomerByEmail :one
SELECT id, name, email, phone, billing_address, created_at FROM customers 
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetCustomerByEmail(ctx context.Context, email pgtype.Text) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomerByEmail, email)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.BillingAddress,
		&i.CreatedAt,
	)
	return i, err
}
