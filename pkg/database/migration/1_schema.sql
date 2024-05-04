CREATE TABLE "users" (
  "id_user" SERIAL PRIMARY KEY,
  "email" varchar unique not null,
  "name" varchar not null,
  "password" varchar not null,
  "salt" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

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

CREATE TABLE "match_cats" (
  "id_match" SERIAL PRIMARY KEY,
  "id_user_cat" integer,
  "id_matched_cat" integer,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "approved_at" timestamp NULL DEFAULT NULL
  "rejected_at" timestamp NULL DEFAULT NULL
);

CREATE TABLE "cat_images" (
  "id_image" SERIAL PRIMARY KEY,
  "id_cat" integer,
  "image" varchar
);

ALTER TABLE "cat_images" ADD FOREIGN KEY ("id_cat") REFERENCES "cats" ("id_cat");

ALTER TABLE "cats" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");

ALTER TABLE "match_cats" ADD FOREIGN KEY ("id_user_cat") REFERENCES "cats" ("id_cat");

ALTER TABLE "match_cats" ADD FOREIGN KEY ("id_matched_cat") REFERENCES "cats" ("id_cat");
