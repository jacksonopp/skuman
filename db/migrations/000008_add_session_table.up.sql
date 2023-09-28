CREATE TABLE "sessions" (
  "id" bigserial PRIMARY KEY,
  "session_id" varchar UNIQUE NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "expires_at" timestamp NOT NULL
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "sessions" ("session_id");