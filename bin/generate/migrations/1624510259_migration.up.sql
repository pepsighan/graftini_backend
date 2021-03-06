BEGIN;
CREATE TABLE IF NOT EXISTS "deployments"("id" uuid NOT NULL, "status" varchar NOT NULL, "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP, "project_deployments" uuid NULL, PRIMARY KEY("id"));
CREATE TABLE IF NOT EXISTS "graph_ql_queries"("id" uuid NOT NULL, "variable_name" varchar NOT NULL, "gql_ast" varchar NOT NULL, "created_at" timestamp with time zone NOT NULL, "updated_at" timestamp with time zone NOT NULL, "project_queries" uuid NULL, PRIMARY KEY("id"));
CREATE TABLE IF NOT EXISTS "pages"("id" uuid NOT NULL, "name" varchar NOT NULL, "route" varchar NOT NULL, "component_map" varchar NOT NULL, "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP, "project_pages" uuid NULL, PRIMARY KEY("id"));
CREATE TABLE IF NOT EXISTS "projects"("id" uuid NOT NULL, "name" varchar NOT NULL, "graphql_endpoint" varchar NULL, "created_at" timestamp with time zone NOT NULL, "updated_at" timestamp with time zone NOT NULL, "user_projects" uuid NULL, PRIMARY KEY("id"));
CREATE TABLE IF NOT EXISTS "users"("id" uuid NOT NULL, "firebase_uid" varchar UNIQUE NOT NULL, "first_name" varchar NULL, "last_name" varchar NULL, "email" varchar NOT NULL, "created_at" timestamp with time zone NOT NULL, "updated_at" timestamp with time zone NOT NULL, PRIMARY KEY("id"));
ALTER TABLE "deployments" ADD CONSTRAINT "deployments_projects_deployments" FOREIGN KEY("project_deployments") REFERENCES "projects"("id") ON DELETE CASCADE;
ALTER TABLE "graph_ql_queries" ADD CONSTRAINT "graph_ql_queries_projects_queries" FOREIGN KEY("project_queries") REFERENCES "projects"("id") ON DELETE CASCADE;
ALTER TABLE "pages" ADD CONSTRAINT "pages_projects_pages" FOREIGN KEY("project_pages") REFERENCES "projects"("id") ON DELETE CASCADE;
ALTER TABLE "projects" ADD CONSTRAINT "projects_users_projects" FOREIGN KEY("user_projects") REFERENCES "users"("id") ON DELETE CASCADE;
COMMIT;
