ALTER TABLE users
ALTER COLUMN "username" DROP NOT NULL,
ADD COLUMN "email" varchar NOT NULL,
ADD COLUMN "verified" boolean;

CREATE INDEX ON "users" ("email");

UPDATE users
SET verified = false;

ALTER TABLE users
ALTER COLUMN "verified" SET NOT NULL;