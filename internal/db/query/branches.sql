-- name: CreateBranch :one
INSERT INTO branches (
  name,
  address,
  organization_id,

) VALUES (
  $1, $2
) RETURNING *;

-- name: GetBranchByID :one
SELECT * from branches 
where id = $ LIMIT 1;