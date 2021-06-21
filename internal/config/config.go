package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment string

var (
	EnvLocal       Environment = "local"
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)

// RequireEnv panics if the environment is not present.
func RequireEnv(key string) string {
	// Load the environment variables from the .env file if it exists.
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("error loading .env file")
	}

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("could not find `%v` environment variable", key)
	}

	return value
}
