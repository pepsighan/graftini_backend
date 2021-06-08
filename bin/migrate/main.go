package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pepsighan/graftini_backend/internal/config"
)

// Run the migrations that are generated in the generate directory.
func main() {
	client, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()

	migrations, err := os.ReadFile("./bin/generate/migration.sql")
	if err != nil {
		log.Fatalf("could not read migration file: %v", err)
	}

	_, err = client.Exec(string(migrations[:]))
	if err != nil {
		log.Fatalf("could not migrate sql: %v", err)
	}
}
