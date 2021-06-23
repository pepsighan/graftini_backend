package config

import (
	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

// Load from the .deploy.env if present.
var _ = config.LoadEnvFile(".deploy.env")

var (
	Env          config.Environment = config.Environment(config.RequireEnv("ENV"))
	VercelToken  string             = config.RequireEnv("VERCEL_TOKEN")
	VercelTeamID string             = config.RequireEnv("VERCEL_TEAM_ID")
	Port         string             = config.Env("PORT", "8888")
)
