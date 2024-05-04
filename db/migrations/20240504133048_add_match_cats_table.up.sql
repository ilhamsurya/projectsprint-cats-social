CREATE TABLE "match_cats" (
  "id_match" SERIAL PRIMARY KEY,
  "id_user_cat" integer,
  "id_matched_cat" integer,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "approved_at" timestamp NULL DEFAULT NULL,
  "rejected_at" timestamp NULL DEFAULT NULL
);

ALTER TABLE "match_cats" ADD FOREIGN KEY ("id_user_cat") REFERENCES "cats" ("id_cat");

ALTER TABLE "match_cats" ADD FOREIGN KEY ("id_matched_cat") REFERENCES "cats" ("id_cat");
