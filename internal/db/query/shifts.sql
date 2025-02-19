-- name: CreateShift :one
INSERT INTO shifts (
    id, staff_id, shift_date, start_time, end_time, status, notes
)
VALUES (uuid_generate_v4(),$1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetShift :one
SELECT 
    shifts.*,
    staff.first_name AS staff_first_name,
    staff.last_name AS staff_last_name
FROM shifts
LEFT JOIN staff ON shifts.staff_id = staff.id
WHERE shifts.id = $1;

-- name: GetAllShifts :one
SELECT 
    shifts.*,
    staff.first_name AS staff_first_name,
    staff.last_name AS staff_last_name
FROM shifts
LEFT JOIN staff ON shifts.staff_id = staff.id;

-- name: UpdateShift :one
UPDATE shifts
SET 
    staff_id = $2, shift_date = $3, start_time = $4, end_time = $5, status = $6, notes = $7
WHERE id = $1
RETURNING *;

-- name: CountShifts :one
SELECT COUNT(*) FROM shifts;