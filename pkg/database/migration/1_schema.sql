CREATE TABLE "users" (
  "id_user" SERIAL PRIMARY KEY,
  "email" varchar unique not null,
  "name" varchar not null,
  "password" varchar not null,
  "salt" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "user_cat" (
  "id_user" integer,
  "id_cat" integer
);

CREATE TABLE "cat" (
  "id_cat" SERIAL PRIMARY KEY,
  "name" varchar,
  "race" varchar,
  "sex" varchar,
  "age_in_month" int,
  "description" varchar,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "match_cat" (
  "id_match" SERIAL PRIMARY KEY,
  "id_user_cat" integer,
  "id_matched_Cat" integer,
  "is_matched" bool
);

CREATE TABLE "cat_image" (
  "id_image" SERIAL PRIMARY KEY,
  "id_cat" integer,
  "image" varchar
);

ALTER TABLE "user_cat" ADD FOREIGN KEY ("id_user") REFERENCES "user" ("id_user");

ALTER TABLE "user_cat" ADD FOREIGN KEY ("id_cat") REFERENCES "cat" ("id_cat");

ALTER TABLE "cat_image" ADD FOREIGN KEY ("id_cat") REFERENCES "cat" ("id_cat");
