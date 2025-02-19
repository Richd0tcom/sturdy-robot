// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: activity_logs.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createActivityLog = `-- name: CreateActivityLog :one
INSERT INTO activity_logs (id, entity_type, entity_id, action, changes, user_id) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5) 
RETURNING id, entity_type, entity_id, action, changes, created_at, user_id
`

type CreateActivityLogParams struct {
	EntityType string      `json:"entity_type"`
	EntityID   pgtype.UUID `json:"entity_id"`
	Action     string      `json:"action"`
	Changes    []byte      `json:"changes"`
	UserID     pgtype.UUID `json:"user_id"`
}

func (q *Queries) CreateActivityLog(ctx context.Context, arg CreateActivityLogParams) (ActivityLog, error) {
	row := q.db.QueryRow(ctx, createActivityLog,
		arg.EntityType,
		arg.EntityID,
		arg.Action,
		arg.Changes,
		arg.UserID,
	)
	var i ActivityLog
	err := row.Scan(
		&i.ID,
		&i.EntityType,
		&i.EntityID,
		&i.Action,
		&i.Changes,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const getActivityLogByEntityID = `-- name: GetActivityLogByEntityID :many
SELECT id, entity_type, entity_id, action, changes, created_at, user_id FROM activity_logs
WHERE entity_id = $1
`

func (q *Queries) GetActivityLogByEntityID(ctx context.Context, entityID pgtype.UUID) ([]ActivityLog, error) {
	rows, err := q.db.Query(ctx, getActivityLogByEntityID, entityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ActivityLog{}
	for rows.Next() {
		var i ActivityLog
		if err := rows.Scan(
			&i.ID,
			&i.EntityType,
			&i.EntityID,
			&i.Action,
			&i.Changes,
			&i.CreatedAt,
			&i.UserID,
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

const getActivityLogsByUserID = `-- name: GetActivityLogsByUserID :many
SELECT id, entity_type, entity_id, action, changes, created_at, user_id FROM activity_logs
WHERE user_id = $1
`

func (q *Queries) GetActivityLogsByUserID(ctx context.Context, userID pgtype.UUID) ([]ActivityLog, error) {
	rows, err := q.db.Query(ctx, getActivityLogsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ActivityLog{}
	for rows.Next() {
		var i ActivityLog
		if err := rows.Scan(
			&i.ID,
			&i.EntityType,
			&i.EntityID,
			&i.Action,
			&i.Changes,
			&i.CreatedAt,
			&i.UserID,
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
