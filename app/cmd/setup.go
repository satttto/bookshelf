package main

import (
	"fmt"

	"github.com/pkg/errors"

	adapterCache "github.com/satttto/bookshelf/app/adapter/cache"
	adapterDB "github.com/satttto/bookshelf/app/adapter/rdb"
	"github.com/satttto/bookshelf/app/config"
	externalDB "github.com/satttto/bookshelf/app/external/postgre"
	externalCache "github.com/satttto/bookshelf/app/external/redis"
	"github.com/satttto/bookshelf/app/pkg/logger"
)

// setupLogger creates a new logger given the log level
func setupLogger(config *config.Config) (*logger.Logger, error) {
	log, err := logger.New(config.LogLevel, logger.EnvType(config.Env))
	if err != nil {
		return nil, err
	}

	return log, nil
}

func setupDB(config *config.Config) (adapterDB.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := externalDB.New(dsn)
	if err != nil {
		return nil, err
	}

	if config.ShouldMigrate {
		if err := db.Migrate(); err != nil {
			return nil, errors.New("Failed to migrate")
		}
	}

	return db, nil
}

func setupCache(config *config.Config) (adapterCache.Cache, error) {
	cache, err := externalCache.New("", config.CacheUser, config.CachePassword)
	if err != nil {
		return nil, err
	}

	return cache, nil
}
