CREATE TABLE "art_categories" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(100) NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "artists" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(200) NOT NULL,
  "biography" TEXT,
  "birth_date" DATE,
  "death_date" DATE,
  "nationality" VARCHAR(100),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "artworks" (
  "id" UUID PRIMARY KEY,
  "title" VARCHAR(200) NOT NULL,
  "artist_id" UUID,
  "category_id" UUID,
  "year_created" INTEGER,
  "medium" VARCHAR(100),
  "dimensions" VARCHAR(100),
  "description" TEXT,
  "acquisition_date" DATE,
  "condition_status" VARCHAR(50),
  "location_in_museum" VARCHAR(100),
  "image_url" VARCHAR(500),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "staff_roles" (
  "id" UUID PRIMARY KEY,
  "title" VARCHAR(100) NOT NULL,
  "description" TEXT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "staff" (
  "id" UUID PRIMARY KEY,
  "first_name" VARCHAR(100) NOT NULL,
  "last_name" VARCHAR(100) NOT NULL,
  "role_id" UUID,
  "email" VARCHAR(200) UNIQUE NOT NULL,
  "phone" VARCHAR(20),
  "hire_date" DATE,
  "status" VARCHAR(50),
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE "shifts" (
  "id" UUID PRIMARY KEY,
  "staff_id" UUID,
  "shift_date" DATE NOT NULL,
  "start_time" TIME NOT NULL,
  "end_time" TIME NOT NULL,
  "status" VARCHAR(50),
  "notes" TEXT,
  "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);

CREATE INDEX "idx_artworks_artist" ON "artworks" ("artist_id");

CREATE INDEX "idx_artworks_category" ON "artworks" ("category_id");

CREATE INDEX "idx_staff_role" ON "staff" ("role_id");

CREATE INDEX "idx_shifts_staff" ON "shifts" ("staff_id");

CREATE INDEX "idx_shifts_date" ON "shifts" ("shift_date");

ALTER TABLE "artworks" ADD FOREIGN KEY ("artist_id") REFERENCES "artists" ("id");

ALTER TABLE "artworks" ADD FOREIGN KEY ("category_id") REFERENCES "art_categories" ("id");

ALTER TABLE "staff" ADD FOREIGN KEY ("role_id") REFERENCES "staff_roles" ("id");

ALTER TABLE "shifts" ADD FOREIGN KEY ("staff_id") REFERENCES "staff" ("id");
