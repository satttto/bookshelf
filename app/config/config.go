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
