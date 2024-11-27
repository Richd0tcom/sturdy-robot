-- name: CreateBranch :one
INSERT INTO branches (
  id, 
  name,
  address,
  organization_id

) VALUES (
  uuid_generate_v4(),
  $1, $2, $3
) RETURNING *;

-- name: GetBranchByID :one
SELECT * from branches 
where id = $1 LIMIT 1;