-- name: CreateStaffRole :one
INSERT INTO staff_roles (id, title, description)
VALUES (uuid_generate_v4(),$1, $2)
RETURNING *;

-- name: GetStaffRole :one
SELECT * FROM staff_roles
WHERE id = $1;

-- name: UpdateStaffRole :one
UPDATE staff_roles
SET title = $2, description = $3
WHERE id = $1
RETURNING *;