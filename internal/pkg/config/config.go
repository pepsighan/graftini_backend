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

// LoadEnvFile loads the environment file if present. This function
// needs to run before any environment is read.
func LoadEnvFile(filename string) bool {
	// Load the environment variables from the .env file if it exists.
	err := godotenv.Load(filename)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("error loading %v file", filename)
	}

	// This value is just a dummy.
	return true
}

// RequireEnv gets the value of the env. It panics if the env does not exist.
func RequireEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("could not find `%v` environment variable", key)
	}

	return value
}
