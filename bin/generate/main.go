package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	unixTime := time.Now().Unix()
	fileUp := fmt.Sprintf("./bin/generate/migrations/%v_migration.up.sql", unixTime)

	fUp, err := os.Create(fileUp)
	if err != nil {
		log.Fatalf("could not create migration file: %v", err)
	}
	defer fUp.Close()

	ctx := context.Background()
	if err := client.Schema.WriteTo(ctx, fUp); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	// Creating a downfile just because its required by migrate package.
	fileDown := fmt.Sprintf("./bin/generate/migrations/%v_migration.down.sql", unixTime)
	fDown, err := os.Create(fileDown)
	if err != nil {
		log.Fatalf("could not create migration file: %v", err)
	}
	defer fDown.Close()
}
