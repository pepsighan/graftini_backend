package backendconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/pepsighan/graftini_backend/internal/config"
)

var (
	Env            config.Environment = config.Environment(config.RequireEnv("ENV"))
	Port                              = port()
	DatabaseURL                       = databaseURL()
	AllowedOrigins                    = allowedOrigins()
	MaxBodySize                       = maxBodySize()
)

func port() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return "1323"
	}
	return port
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
