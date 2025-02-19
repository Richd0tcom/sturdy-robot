-- Drop foreign key constraints
ALTER TABLE "shifts" DROP CONSTRAINT IF EXISTS "shifts_staff_id_fkey";
ALTER TABLE "staff" DROP CONSTRAINT IF EXISTS "staff_role_id_fkey";
ALTER TABLE "artworks" DROP CONSTRAINT IF EXISTS "artworks_category_id_fkey";
ALTER TABLE "artworks" DROP CONSTRAINT IF EXISTS "artworks_artist_id_fkey";

-- Drop indexes
DROP INDEX IF EXISTS "idx_shifts_date";
DROP INDEX IF EXISTS "idx_shifts_staff";
DROP INDEX IF EXISTS "idx_staff_role";
DROP INDEX IF EXISTS "idx_artworks_category";
DROP INDEX IF EXISTS "idx_artworks_artist";

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS "shifts";
DROP TABLE IF EXISTS "staff";
DROP TABLE IF EXISTS "staff_roles";
DROP TABLE IF EXISTS "artworks";
DROP TABLE IF EXISTS "artists";
DROP TABLE IF EXISTS "art_categories";