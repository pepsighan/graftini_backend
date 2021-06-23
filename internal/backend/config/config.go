package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

// Load from the .backend.env if present.
var _ = config.LoadEnvFile(".backend.env")

var (
	Env            config.Environment = config.Environment(config.RequireEnv("ENV"))
	Port                              = config.Env("PORT", "1323")
	DatabaseURL                       = databaseURL()
	AllowedOrigins                    = allowedOrigins()
	MaxBodySize                       = config.Env("MAX_BODY_SIZE", "2M")
	DeployEndpoint                    = config.RequireEnv("DEPLOY_ENDPOINT")
)

func allowedOrigins() []string {
	allowedOrigins, ok := os.LookupEnv("ALLOWED_ORIGINS")
	if ok {
		return strings.Split(allowedOrigins, ",")
	} else {
		return []string{}
	}
}

// databaseURL gets the database URL to connect to based on the environment.
func databaseURL() string {
	if Env == config.EnvLocal {
		return config.RequireEnv("POSTGRES_URL")
	}

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dbUser := config.RequireEnv("DB_USER")
	dbPassword := config.RequireEnv("DB_PASS")
	dbInstanceConnectionName := config.RequireEnv("INSTANCE_CONNECTION_NAME")
	dbName := config.RequireEnv("DB_NAME")

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPassword, dbName, socketDir, dbInstanceConnectionName)
	return dbURI
}