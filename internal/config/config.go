package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	Env                          Environment = Environment(requireEnv("ENV"))
	DatabaseURL                              = databaseURL()
	AllowedOrigins                           = allowedOrigins()
	MaxBodySize                              = maxBodySize()
	GoogleApplicationCredentials             = requireEnv("GOOGLE_APPLICATION_CREDENTIALS")
)

type Environment string

var (
	EnvLocal       Environment = "local"
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)

// requireEnv panics if the environment is not present.
func requireEnv(key string) string {
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

func allowedOrigins() []string {
	allowedOrigins, ok := os.LookupEnv("ALLOWED_ORIGINS")
	if ok {
		return strings.Split(allowedOrigins, ",")
	} else {
		return []string{}
	}
}

func maxBodySize() string {
	maxBodySize, ok := os.LookupEnv("MAX_BODY_SIZE")
	if ok {
		return maxBodySize
	} else {
		return "2M"
	}
}

// databaseURL gets the database URL to connect to based on the environment.
func databaseURL() string {
	if Env == EnvLocal {
		return requireEnv("POSTGRES_URL")
	}

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dbUser := requireEnv("DB_USER")
	dbPassword := requireEnv("DB_PASS")
	dbInstanceConnectionName := requireEnv("INSTANCE_CONNECTION_NAME")
	dbName := requireEnv("DB_NAME")

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPassword, dbName, socketDir, dbInstanceConnectionName)
	return dbURI
}
