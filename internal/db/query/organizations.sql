-- name: CreateOrganization :one
INSERT INTO organizations (
  name,
  email

) VALUES (
  $1, $2
) RETURNING *;

-- name: GetOrganizationByID :one
SELECT * from users 
where id = $1 LIMIT 1;