CREATE TABLE "cats" (
  "id_cat" SERIAL PRIMARY KEY,
  "id_user" integer,
  "name" varchar,
  "race" varchar,
  "sex" varchar,
  "age_in_month" int,
  "description" varchar,
  "has_matched" bool DEFAULT FALSE,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp NULL DEFAULT NULL
);

CREATE TABLE "cat_images" (
  "id_image" SERIAL PRIMARY KEY,
  "id_cat" integer,
  "image" varchar
);

ALTER TABLE "cats" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");

ALTER TABLE "cat_images" ADD FOREIGN KEY ("id_cat") REFERENCES "cats" ("id_cat");