CREATE TABLE "users" (
  "id_user" SERIAL PRIMARY KEY,
  "email" varchar unique not null,
  "name" varchar not null,
  "password" varchar not null,
  "salt" varchar not null,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);