-- name: CreateArtCategory :one
INSERT INTO art_categories (id, name, description)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;

-- name: GetArtCategory :one
SELECT * FROM art_categories
WHERE id = $1;

-- name: GetAllArtCategories :many
SELECT * FROM art_categories;

-- name: UpdateArtCategory :one
UPDATE art_categories
SET name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteArtCategory :exec
DELETE FROM art_categories
WHERE id = $1;

-- name: CountArtCategories :one
SELECT COUNT(*) FROM art_categories;