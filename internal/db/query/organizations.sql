-- name: CreateOrganization :one
INSERT INTO organizations (
  id,
  name,
  email

) VALUES (
  uuid_generate_v4(),
  $1, $2
) RETURNING *;

-- name: GetOrganizationByID :one
SELECT * from organizations 
where id = $1 LIMIT 1;