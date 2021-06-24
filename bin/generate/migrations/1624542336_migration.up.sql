BEGIN;
ALTER TABLE "deployments" ADD COLUMN "vercel_deployment_id" varchar NOT NULL DEFAULT '';
COMMIT;
