BEGIN;
ALTER TABLE "users" ADD COLUMN "is_admin" boolean NULL;
COMMIT;
