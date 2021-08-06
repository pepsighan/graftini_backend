BEGIN;
ALTER TABLE "templates" ADD COLUMN "preview_file_id" uuid NULL;
COMMIT;
