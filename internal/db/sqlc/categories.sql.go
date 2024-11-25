// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: categories.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (id, name, description, parent_id, branch_id) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4) 
RETURNING id, parent_id, name, branch_id, description, created_at
`

type CreateCategoryParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ParentID    uuid.NullUUID  `json:"parent_id"`
	BranchID    uuid.UUID      `json:"branch_id"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.Name,
		arg.Description,
		arg.ParentID,
		arg.BranchID,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.ParentID,
		&i.Name,
		&i.BranchID,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const getCategoriesByBranchID = `-- name: GetCategoriesByBranchID :many
SELECT id, parent_id, name, branch_id, description, created_at FROM categories 
WHERE branch_id = $1
`

func (q *Queries) GetCategoriesByBranchID(ctx context.Context, branchID uuid.UUID) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategoriesByBranchID, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Name,
			&i.BranchID,
			&i.Description,
			&i.CreatedAt,
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