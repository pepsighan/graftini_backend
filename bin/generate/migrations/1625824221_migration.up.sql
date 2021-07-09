BEGIN;
CREATE TABLE IF NOT EXISTS "files"("id" uuid NOT NULL, "kind" varchar NOT NULL, "mime_type" varchar NOT NULL, "created_at" timestamp with time zone NOT NULL, "updated_at" timestamp with time zone NOT NULL, PRIMARY KEY("id"));
COMMIT;
