package main

import (
	"github.com/satttto/bookshelf/app/config"
	"github.com/satttto/bookshelf/app/pkg/logger"
)

// setupLogger creates a new logger given the log level
func setupLogger(config *config.Config) (*logger.Log, error) {
	log, err := logger.New(config.LogLevel, logger.EnvType(config.Env))
	if err != nil {
		return nil, err
	}

	return log, nil
}
