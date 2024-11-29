// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: currency.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCurrency = `-- name: CreateCurrency :one
INSERT INTO currency (id, name, symbol, code) 
VALUES (uuid_generate_v4(), $1, $2, $3) 
RETURNING id, name, code, symbol, created_at
`

type CreateCurrencyParams struct {
	Name   string      `json:"name"`
	Symbol pgtype.Text `json:"symbol"`
	Code   string      `json:"code"`
}

func (q *Queries) CreateCurrency(ctx context.Context, arg CreateCurrencyParams) (Currency, error) {
	row := q.db.QueryRow(ctx, createCurrency, arg.Name, arg.Symbol, arg.Code)
	var i Currency
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Symbol,
		&i.CreatedAt,
	)
	return i, err
}

const getCurrencyByID = `-- name: GetCurrencyByID :one
SELECT id, name, code, symbol, created_at FROM currency 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCurrencyByID(ctx context.Context, id pgtype.UUID) (Currency, error) {
	row := q.db.QueryRow(ctx, getCurrencyByID, id)
	var i Currency
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Symbol,
		&i.CreatedAt,
	)
	return i, err
}
