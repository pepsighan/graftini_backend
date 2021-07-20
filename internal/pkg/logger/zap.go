package logger

import (
	"log"

	"github.com/blendle/zapdriver"
	"github.com/pepsighan/graftini_backend/internal/pkg/config"
	"go.uber.org/zap"
)

// NewLogger creates a new logger and replaces the global zap logger with it.
// Do not forget do `defer logger.Sync()` in main to flush any logs on exit.
func NewLogger(env config.Environment) *zap.Logger {
	var logger *zap.Logger
	var err error

	if env.IsProduction() {
		logger, err = zapdriver.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("could not create a logger instance: %v", err)
	}

	// Using this logger instance when using global logger.
	zap.ReplaceGlobals(logger)
	return logger
}
