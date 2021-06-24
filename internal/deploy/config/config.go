package config

import (
	"strings"

	"github.com/pepsighan/graftini_backend/internal/pkg/config"
)

// Load from the .deploy.env if present.
var _ = config.LoadEnvFile(".deploy.env")

var (
	Env                 = config.Environment(config.RequireEnv("ENV"))
	DatabaseURL         = config.DatabaseURL(Env)
	VercelToken         = config.RequireEnv("VERCEL_TOKEN")
	VercelTeamID        = config.RequireEnv("VERCEL_TEAM_ID")
	Port                = config.Env("PORT", "8888")
	TemplateNextAppPath = config.RequireEnv("TEMPLATE_NEXT_APP_PATH")
	GitHubNPMRepoToken  = strings.Replace(config.RequireEnv("GITHUB_NPM_REPO_TOKEN"), "\\n", "\n", 1) // Replace \n with actual new line.
)
