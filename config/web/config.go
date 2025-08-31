package web

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App      App
	Api      Api
	Keycloak Keycloak
	Session  Session
}

type App struct {
	Name    string `env:"APP_NAME"`
	DevMode bool   `env:"APP_DEV_MODE" envDefault:"false"`
}

type Api struct {
	Host string `env:"API_HOST"`
}

type Keycloak struct {
	UIHost       string `env:"KEYCLOAK_UI_HOST"`
	APIHost      string `env:"KEYCLOAK_API_HOST"`
	ClientId     string `env:"KEYCLOAK_OAUTH_CLIENT_ID"`
	ClientSecret string `env:"KEYCLOAK_OAUTH_CLIENT_SECRET"`
	Realm        string `env:"KEYCLOAK_REALM" envDefault:"app"`
	RedirectURI  string `env:"KEYCLOAK_REDIRECT_URI"`
}

type Session struct {
	Driver         string `env:"SESSION_DRIVER" envDefault:"cookie"`
	SessionSecret  string `env:"SESSION_SECRET"`
	SessionKey     string `env:"SESSION_KEY"`
	RedisPoolSize  int    `env:"REDIS_POOL_SIZE" envDefault:"10"`
	RedisNetwork   string `env:"REDIS_NETWORK" envDefault:"tcp"`
	RedisAddr      string `env:"REDIS_ADDR" envDefault:"redis:6379"`
	RedisUsername  string `env:"REDIS_USERNAME"`
	RedisPassword  string `env:"REDIS_PASSWORD"`
	FilesystemPath string `env:"FILESYSTEM_PATH" envDefault:"/tmp/go-apps/"`
}

// NewConfig returns app config.
func NewConfig() (Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return *cfg, fmt.Errorf("config error: %w", err)
	}

	return *cfg, nil
}
