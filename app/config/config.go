package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	envLocal       = "local"
	envDevelopment = "dev"
	envProduction  = "prd"
)

type Config struct {
	Env      string `envconfig:"ENV" required:"true"`
	GRPCPort int    `envconfig:"GRPCPort" default:"8080"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	// Database config
	DBName        string `envconfig:"DB_NAME" default:"bookshelf"`
	DBUser        string `envconfig:"DB_USER" default:"bookshelf-user"`
	DBPassword    string `envconfig:"DB_PASSWORD" required:"true"`
	DBPort        string `envconfig:"DB_PORT" default:"5432"`
	DBHost        string `envconfig:"DB_HOST" default:"localhost"`
	ShouldMigrate bool   `envconfig:"SHOULD_MIGRATE" default:"false"`
}

func ReadConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) IsProduction() bool {
	return c.Env == envProduction
}
