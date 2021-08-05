BEGIN;
CREATE TABLE IF NOT EXISTS "templates"("id" uuid NOT NULL, "name" varchar NOT NULL, "snapshot" varchar NOT NULL, "created_at" timestamp with time zone NOT NULL, "updated_at" timestamp with time zone NOT NULL, PRIMARY KEY("id"));
COMMIT;
