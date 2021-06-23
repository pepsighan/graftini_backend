package config

import (
	"os"
	"strings"

	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

// Load from the .backend.env if present.
var _ = config.LoadEnvFile(".backend.env")

var (
	Env            = config.Environment(config.RequireEnv("ENV"))
	Port           = config.Env("PORT", "1323")
	DatabaseURL    = config.DatabaseURL(Env)
	AllowedOrigins = allowedOrigins()
	MaxBodySize    = config.Env("MAX_BODY_SIZE", "2M")
	DeployEndpoint = config.RequireEnv("DEPLOY_ENDPOINT")
)

func allowedOrigins() []string {
	allowedOrigins, ok := os.LookupEnv("ALLOWED_ORIGINS")
	if ok {
		return strings.Split(allowedOrigins, ",")
	} else {
		return []string{}
	}
}
