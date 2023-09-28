CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "verified" boolean NOT NULL,
  "verification_code" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" bigserial PRIMARY KEY,
  "session_id" varchar UNIQUE NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "expires_at" timestamp NOT NULL
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "sessions" ("session_id");
