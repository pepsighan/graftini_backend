BEGIN;
CREATE TABLE IF NOT EXISTS "early_accesses"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, "email" varchar UNIQUE NOT NULL, "created_at" timestamp with time zone NOT NULL, "updated_at" timestamp with time zone NOT NULL, PRIMARY KEY("id"));
COMMIT;