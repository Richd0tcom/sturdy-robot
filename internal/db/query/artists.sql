-- name: CreateArtist :one
INSERT INTO artists (id, name, biography, birth_date, death_date, nationality)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)
RETURNING *;


-- name: GetArtist :one
SELECT * FROM artists
WHERE id = $1;

-- name: GetAllArtists :many
SELECT * FROM artists;

-- name: UpdateArtist :one
UPDATE artists
SET name = $2, biography = $3, birth_date = $4, death_date = $5, nationality = $6
WHERE id = $1
RETURNING *;

-- name: CountArtists :one
SELECT COUNT(*) FROM artists;

