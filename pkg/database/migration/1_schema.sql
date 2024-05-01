CREATE TYPE "race_list" AS ENUM (
  'Persian',
  'Maine_Coon',
  'Siamese',
  'Ragdoll',
  'Bengal',
  'Sphynx',
  'British_Shorthair',
  'Abyssinian',
  'Scottish_Fold',
  'Birman'
);

CREATE TYPE "sex_list" AS ENUM (
  'male',
  'female'
);

CREATE TABLE "users" (
  "id_user" serial PRIMARY KEY,
  "email" varchar unique not null,
  "name" varchar not null,
  "password" varchar not null,
  "created_at" timestamp default current_timestamp,
  "updated_at" timestamp
);

CREATE TABLE "user_cat" (
  "id_user" integer,
  "id_cat" integer
);

CREATE TABLE "cat" (
  "id_cat" integer PRIMARY KEY,
  "name" varchar,
  "race" race_list,
  "sex" sex_list,
  "age_in_month" int,
  "description" varchar,
  "image_url" varchar,
  "isMatch" bool,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "match_cat" (
  "id_match" int PRIMARY KEY,
  "id_user_cat" int,
  "id_matched_Cat" int
);

CREATE TABLE "match_request" (
  "id_request" int PRIMARY KEY,
  "id_user_cat" int,
  "id_matched_Cat" int
);

ALTER TABLE "user_cat" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id_user");

ALTER TABLE "user_cat" ADD FOREIGN KEY ("id_cat") REFERENCES "cat" ("id_cat");
