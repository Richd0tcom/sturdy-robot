-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  address,
  branch_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserById :one
SELECT * from users 
where id = $1 LIMIT 1;