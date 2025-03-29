package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	HttpPort string `env:"HTTP_PORT" envDefault:"8080"`
	Env      string `env:"ENV" envDefault:"development"`
	AppName  string `env:"APP_NAME"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env vars: %w", err)
	}

	return cfg, nil
}
