package main

import (
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pepsighan/graftini_backend/ent"
	"github.com/pepsighan/graftini_backend/internal/config"
)

// Generate the migration manually. Do not automigrate. We can track
// if any issues exist with the generated SQL.
func main() {
	client, err := ent.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()

	f, err := os.Create("./bin/migrate/migration.sql")
	if err != nil {
		log.Fatalf("could not create migration file: %v", err)
	}
	defer f.Close()

	ctx := context.Background()
	if err := client.Schema.WriteTo(ctx, f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
