// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: organizations.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO organizations (
  id,
  name,
  email

) VALUES (
  uuid_generate_v4(),
  $1, $2
) RETURNING id, name, email, active, created_at
`

type CreateOrganizationParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error) {
	row := q.db.QueryRow(ctx, createOrganization, arg.Name, arg.Email)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Active,
		&i.CreatedAt,
	)
	return i, err
}

const getOrganizationByID = `-- name: GetOrganizationByID :one
SELECT id, name, email, active, created_at from organizations 
where id = $1 LIMIT 1
`

func (q *Queries) GetOrganizationByID(ctx context.Context, id pgtype.UUID) (Organization, error) {
	row := q.db.QueryRow(ctx, getOrganizationByID, id)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Active,
		&i.CreatedAt,
	)
	return i, err
}
