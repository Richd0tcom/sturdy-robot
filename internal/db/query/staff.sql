-- name: CreateStaff :one
INSERT INTO staff (
    id, first_name, last_name, role_id, email, phone, hire_date, status
)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetStaff :one
SELECT 
    staff.*,
    staff_roles.title AS role_title
FROM staff
LEFT JOIN staff_roles ON staff.role_id = staff_roles.id
WHERE staff.id = $1;

-- name: GetAllStaff :many
SELECT 
    staff.*,
    staff_roles.title AS role_title
FROM staff
LEFT JOIN staff_roles ON staff.role_id = staff_roles.id;

-- name: UpdateStaff :one
UPDATE staff
SET 
    first_name = $2, last_name = $3, role_id = $4, email = $5, phone = $6,
    hire_date = $7, status = $8
WHERE id = $1
RETURNING *;

-- name: DeleteStaff :exec
DELETE FROM staff
WHERE id = $1;

-- name: CountStaff :one
SELECT COUNT(*) FROM staff;