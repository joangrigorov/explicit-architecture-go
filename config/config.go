package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App      App
	DB       DB
	Tracing  Tracing
	Keycloak Keycloak
}

type App struct {
	Name    string `env:"APP_NAME"`
	DevMode bool   `env:"APP_DEV_MODE" envDefault:"false"`
}
type DB struct {
	Activity   DBActivity
	Attendance DBAttendance
	User       DBUser
}

type DBActivity struct {
	Driver          string `env:"DB_ACTIVITY_DRIVER" envDefault:"postgres"`
	Host            string `env:"DB_ACTIVITY_HOST" envDefault:"localhost"`
	Port            string `env:"DB_ACTIVITY_PORT" envDefault:"5432"`
	User            string `env:"DB_ACTIVITY_USER" envDefault:"postgres"`
	Password        string `env:"DB_ACTIVITY_PASSWORD" envDefault:"secret"`
	Database        string `env:"DB_ACTIVITY_DATABASE" envDefault:"activity"`
	SslMode         string `env:"DB_ACTIVITY_SSLMODE" envDefault:"disable"`
	MaxOpenConns    int    `env:"DB_ACTIVITY_MAX_OPEN_CONNS" envDefault:"5"`
	MaxIdleConns    int    `env:"DB_ACTIVITY_MAX_IDLE_CONNS" envDefault:"2"`
	ConnMaxLifetime int    `env:"DB_ACTIVITY_CONN_MAX_LIFETIME_MINS" envDefault:"30"`
}

type DBAttendance struct {
	Driver          string `env:"DB_ATTENDANCE_DRIVER" envDefault:"postgres"`
	Host            string `env:"DB_ATTENDANCE_HOST" envDefault:"localhost"`
	Port            string `env:"DB_ATTENDANCE_PORT" envDefault:"5432"`
	User            string `env:"DB_ATTENDANCE_USER" envDefault:"postgres"`
	Password        string `env:"DB_ATTENDANCE_PASSWORD" envDefault:"secret"`
	Database        string `env:"DB_ATTENDANCE_DATABASE" envDefault:"attendance"`
	SslMode         string `env:"DB_ATTENDANCE_SSLMODE" envDefault:"disable"`
	MaxOpenConns    int    `env:"DB_ATTENDANCE_MAX_OPEN_CONNS" envDefault:"5"`
	MaxIdleConns    int    `env:"DB_ATTENDANCE_MAX_IDLE_CONNS" envDefault:"2"`
	ConnMaxLifetime int    `env:"DB_ATTENDANCE_CONN_MAX_LIFETIME_MINS" envDefault:"30"`
}

type DBUser struct {
	Driver          string `env:"DB_USER_DRIVER" envDefault:"postgres"`
	Host            string `env:"DB_USER_HOST" envDefault:"localhost"`
	Port            string `env:"DB_USER_PORT" envDefault:"5432"`
	User            string `env:"DB_USER_USER" envDefault:"postgres"`
	Password        string `env:"DB_USER_PASSWORD" envDefault:"secret"`
	Database        string `env:"DB_USER_DATABASE" envDefault:"user"`
	SslMode         string `env:"DB_USER_SSLMODE" envDefault:"disable"`
	MaxOpenConns    int    `env:"DB_USER_MAX_OPEN_CONNS" envDefault:"5"`
	MaxIdleConns    int    `env:"DB_USER_MAX_IDLE_CONNS" envDefault:"2"`
	ConnMaxLifetime int    `env:"DB_USER_CONN_MAX_LIFETIME_MINS" envDefault:"30"`
}

type Tracing struct {
	Endpoint string `env:"TRACE_ENDPOINT"`
}

type Keycloak struct {
	Url          string `env:"KEYCLOAK_URL"`
	ClientId     string `env:"KEYCLOAK_OAUTH_CLIENT_ID"`
	ClientSecret string `env:"KEYCLOAK_OAUTH_CLIENT_SECRET"`
	Realm        string `env:"KEYCLOAK_REALM" envDefault:"app"`
	Scopes       string `env:"KEYCLOAK_SCOPES"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
