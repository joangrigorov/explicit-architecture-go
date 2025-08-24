package web

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App      App
	Api      Api
	Keycloak Keycloak
}

type App struct {
	Name    string `env:"APP_NAME"`
	DevMode bool   `env:"APP_DEV_MODE" envDefault:"false"`
}

type Api struct {
	Host string `env:"API_HOST"`
}

type Keycloak struct {
	Host        string `env:"KEYCLOAK_HOST"`
	ClientId    string `env:"KEYCLOAK_OAUTH_CLIENT_ID"`
	Realm       string `env:"KEYCLOAK_REALM" envDefault:"app"`
	RedirectURI string `env:"KEYCLOAK_REDIRECT_URI"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
