-- remove update function
DROP TRIGGER IF EXISTS update_user_modtime ON users;

DROP FUNCTION IF EXISTS update_updated_at_column();

ALTER TABLE "users"
DROP COLUMN "created_at",
DROP COLUMN "updated_at";