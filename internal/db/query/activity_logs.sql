-- name: CreateActivityLog :one
INSERT INTO activity_logs (id, entity_type, entity_id, action, changes, user_id) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetActivityLogsByUserID :many
SELECT * FROM activity_logs
WHERE user_id = $1;

-- name: GetActivityLogByEntityID :many
SELECT * FROM activity_logs
WHERE entity_id = $1;