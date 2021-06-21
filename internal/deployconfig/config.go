package deployconfig

import "github.com/pepsighan/graftini_backend/internal/config"

var (
	Env          Environment = Environment(config.RequireEnv("ENV"))
	VercelToken  string      = config.RequireEnv("VERCEL_TOKEN")
	VercelTeamID string      = config.RequireEnv("VERCEL_TEAM_ID")
)

type Environment string

var (
	EnvLocal       Environment = "local"
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)
