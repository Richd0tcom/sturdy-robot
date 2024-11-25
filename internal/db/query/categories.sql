-- name: CreateCategory :one
INSERT INTO categories (id, name, description, parent_id, branch_id) 
VALUES (uuid_generate_v4(), $1, $2, $3, $4) 
RETURNING *;

-- name: GetCategoriesByBranchID
SELECT * FROM categories 
WHERE branch_id = $1;