package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DB DB
	PG PG
}

type DB struct {
	Driver string `env:"DB_DRIVER" envDefault:"postgres"`
	PG     PG
}

type PG struct {
	PoolMax int    `env:"PG_POOL_MAX,required"`
	URL     string `env:"PG_URL,required"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
