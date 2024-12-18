// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product_versions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProductVersion = `-- name: CreateProductVersion :one
INSERT INTO product_versions (
    id, product_id, branch_id, name, 
    price_adjustment, attributes, stock_quantity, reorder_point
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7
) RETURNING id, product_id, branch_id, name, price_adjustment, attributes, stock_quantity, reorder_point, created_at
`

type CreateProductVersionParams struct {
	ProductID       pgtype.UUID    `json:"product_id"`
	BranchID        pgtype.UUID    `json:"branch_id"`
	Name            string         `json:"name"`
	PriceAdjustment pgtype.Numeric `json:"price_adjustment"`
	Attributes      []byte         `json:"attributes"`
	StockQuantity   pgtype.Int4    `json:"stock_quantity"`
	ReorderPoint    pgtype.Int4    `json:"reorder_point"`
}

func (q *Queries) CreateProductVersion(ctx context.Context, arg CreateProductVersionParams) (ProductVersion, error) {
	row := q.db.QueryRow(ctx, createProductVersion,
		arg.ProductID,
		arg.BranchID,
		arg.Name,
		arg.PriceAdjustment,
		arg.Attributes,
		arg.StockQuantity,
		arg.ReorderPoint,
	)
	var i ProductVersion
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.BranchID,
		&i.Name,
		&i.PriceAdjustment,
		&i.Attributes,
		&i.StockQuantity,
		&i.ReorderPoint,
		&i.CreatedAt,
	)
	return i, err
}

const getProductVersionsByProductID = `-- name: GetProductVersionsByProductID :many
SELECT id, product_id, branch_id, name, price_adjustment, attributes, stock_quantity, reorder_point, created_at FROM product_versions WHERE product_id = $1
`

func (q *Queries) GetProductVersionsByProductID(ctx context.Context, productID pgtype.UUID) ([]ProductVersion, error) {
	rows, err := q.db.Query(ctx, getProductVersionsByProductID, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductVersion{}
	for rows.Next() {
		var i ProductVersion
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.BranchID,
			&i.Name,
			&i.PriceAdjustment,
			&i.Attributes,
			&i.StockQuantity,
			&i.ReorderPoint,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductVersion = `-- name: UpdateProductVersion :one
UPDATE product_versions 
SET 
    name = $2, 
    price_adjustment = $3,
    attributes = $4,
    stock_quantity = $5,
    reorder_point = $6
WHERE id = $1 
RETURNING id, product_id, branch_id, name, price_adjustment, attributes, stock_quantity, reorder_point, created_at
`

type UpdateProductVersionParams struct {
	ID              pgtype.UUID    `json:"id"`
	Name            string         `json:"name"`
	PriceAdjustment pgtype.Numeric `json:"price_adjustment"`
	Attributes      []byte         `json:"attributes"`
	StockQuantity   pgtype.Int4    `json:"stock_quantity"`
	ReorderPoint    pgtype.Int4    `json:"reorder_point"`
}

func (q *Queries) UpdateProductVersion(ctx context.Context, arg UpdateProductVersionParams) (ProductVersion, error) {
	row := q.db.QueryRow(ctx, updateProductVersion,
		arg.ID,
		arg.Name,
		arg.PriceAdjustment,
		arg.Attributes,
		arg.StockQuantity,
		arg.ReorderPoint,
	)
	var i ProductVersion
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.BranchID,
		&i.Name,
		&i.PriceAdjustment,
		&i.Attributes,
		&i.StockQuantity,
		&i.ReorderPoint,
		&i.CreatedAt,
	)
	return i, err
}
