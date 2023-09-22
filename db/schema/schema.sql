CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "verified" boolean NOT NULL,
  "verification_code" varchar
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");
