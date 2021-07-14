BEGIN;
ALTER TABLE "deployments" ADD COLUMN "project_snapshot" varchar NOT NULL DEFAULT '';
COMMIT;
