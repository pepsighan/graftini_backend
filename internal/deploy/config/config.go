package config

import (
	"os"

	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

var (
	Env          Environment = Environment(config.RequireEnv("ENV"))
	VercelToken  string      = config.RequireEnv("VERCEL_TOKEN")
	VercelTeamID string      = config.RequireEnv("VERCEL_TEAM_ID")
	Port         string      = port()
)

type Environment string

var (
	EnvLocal       Environment = "local"
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)

func port() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return "1323"
	}
	return port
}
