// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: payment_info.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPaymentInfo = `-- name: CreatePaymentInfo :one
INSERT INTO payment_info (id, user_id, account_no, routing_no, account_name, bank_name)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5) RETURNING id, user_id, account_no, routing_no, account_name, bank_name, created_at
`

type CreatePaymentInfoParams struct {
	UserID      uuid.NullUUID  `json:"user_id"`
	AccountNo   sql.NullString `json:"account_no"`
	RoutingNo   sql.NullString `json:"routing_no"`
	AccountName sql.NullString `json:"account_name"`
	BankName    sql.NullString `json:"bank_name"`
}

// Insert payment info
func (q *Queries) CreatePaymentInfo(ctx context.Context, arg CreatePaymentInfoParams) (PaymentInfo, error) {
	row := q.db.QueryRowContext(ctx, createPaymentInfo,
		arg.UserID,
		arg.AccountNo,
		arg.RoutingNo,
		arg.AccountName,
		arg.BankName,
	)
	var i PaymentInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccountNo,
		&i.RoutingNo,
		&i.AccountName,
		&i.BankName,
		&i.CreatedAt,
	)
	return i, err
}

const getPaymentInfoByUserID = `-- name: GetPaymentInfoByUserID :one
SELECT id, user_id, account_no, routing_no, account_name, bank_name, created_at FROM payment_info WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetPaymentInfoByUserID(ctx context.Context, userID uuid.NullUUID) (PaymentInfo, error) {
	row := q.db.QueryRowContext(ctx, getPaymentInfoByUserID, userID)
	var i PaymentInfo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccountNo,
		&i.RoutingNo,
		&i.AccountName,
		&i.BankName,
		&i.CreatedAt,
	)
	return i, err
}

const updatePaymentInfo = `-- name: UpdatePaymentInfo :exec
UPDATE payment_info 
SET account_no = $2, routing_no = $3, account_name = $4, bank_name = $5
WHERE user_id = $1
`

type UpdatePaymentInfoParams struct {
	UserID      uuid.NullUUID  `json:"user_id"`
	AccountNo   sql.NullString `json:"account_no"`
	RoutingNo   sql.NullString `json:"routing_no"`
	AccountName sql.NullString `json:"account_name"`
	BankName    sql.NullString `json:"bank_name"`
}

func (q *Queries) UpdatePaymentInfo(ctx context.Context, arg UpdatePaymentInfoParams) error {
	_, err := q.db.ExecContext(ctx, updatePaymentInfo,
		arg.UserID,
		arg.AccountNo,
		arg.RoutingNo,
		arg.AccountName,
		arg.BankName,
	)
	return err
}
