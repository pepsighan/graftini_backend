package config

import (
	"fmt"
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

func (e Environment) IsLocal() bool {
	return e == EnvLocal
}

func (e Environment) IsDevelopment() bool {
	return e == EnvDevelopment
}

func (e Environment) IsProduction() bool {
	return e == EnvProduction
}

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

// Env gets the value from the env. If not present, uses the default value.
func Env(key string, defaultVal string) string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return defaultVal
	}
	return port
}

// DatabaseURL gets the database URL to connect to based on the environment.
func DatabaseURL(env Environment) string {
	if env.IsLocal() {
		return RequireEnv("POSTGRES_URL")
	}

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dbUser := RequireEnv("DB_USER")
	dbPassword := RequireEnv("DB_PASS")
	dbInstanceConnectionName := RequireEnv("INSTANCE_CONNECTION_NAME")
	dbName := RequireEnv("DB_NAME")

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPassword, dbName, socketDir, dbInstanceConnectionName)
	return dbURI
}
