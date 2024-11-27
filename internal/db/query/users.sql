-- name: CreateUser :one
INSERT INTO users (
  id,
  name,
  email,
  address,
  branch_id
) VALUES (
  uuid_generate_v4(),
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserById :one
SELECT * from users 
where id = $1 LIMIT 1;