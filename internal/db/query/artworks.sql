-- name: CreateArtwork :one
INSERT INTO artworks (
    id, title, artist_id, category_id, year_created, medium, dimensions,
    description, acquisition_date, condition_status, location_in_museum, image_url
)
VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetArtwork :one
SELECT 
    artworks.*,
    artists.name AS artist_name,
    art_categories.name AS category_name
FROM artworks
LEFT JOIN artists ON artworks.artist_id = artists.id
LEFT JOIN art_categories ON artworks.category_id = art_categories.id
WHERE artworks.id = $1;

-- name: GetAllArtwork :one
SELECT 
    artworks.*,
    artists.name AS artist_name,
    art_categories.name AS category_name
FROM artworks
LEFT JOIN artists ON artworks.artist_id = artists.id
LEFT JOIN art_categories ON artworks.category_id = art_categories.id;

-- name: UpdateArtwork :one
UPDATE artworks
SET 
    title = $2, artist_id = $3, category_id = $4, year_created = $5, medium = $6,
    dimensions = $7, description = $8, acquisition_date = $9, condition_status = $10,
    location_in_museum = $11, image_url = $12
WHERE id = $1
RETURNING *;

-- name: CountArtworks :one
SELECT COUNT(*) FROM artworks;