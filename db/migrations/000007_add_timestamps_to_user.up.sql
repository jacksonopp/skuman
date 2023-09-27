ALTER TABLE "users"
ADD COLUMN "created_at" timestamp DEFAULT (now()),
ADD COLUMN "updated_at" timestamp DEFAULT (now());

-- create update function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END
$$ language 'plpgsql';

CREATE TRIGGER update_user_modtime BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
