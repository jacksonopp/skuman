CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "verified" boolean NOT NULL,
  "verification_code" varchar
);

CREATE INDEX ON "users" ("email");
