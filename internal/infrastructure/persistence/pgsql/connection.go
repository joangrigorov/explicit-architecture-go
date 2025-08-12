package pgsql

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/lib/pq"
)

type SSLMode string

const (
	SSLModeDisable    SSLMode = "disable"
	SSLModeAllow      SSLMode = "allow"
	SSLModePrefer     SSLMode = "prefer"
	SSLModeRequire    SSLMode = "require"
	SSLModeVerifyCA   SSLMode = "verify-ca"
	SSLModeVerifyFull SSLMode = "verify-full"
)

var validSSLModes = map[SSLMode]struct{}{
	SSLModeDisable:    {},
	SSLModeAllow:      {},
	SSLModePrefer:     {},
	SSLModeRequire:    {},
	SSLModeVerifyCA:   {},
	SSLModeVerifyFull: {},
}

func validateSSLMode(mode SSLMode) {
	if _, ok := validSSLModes[mode]; !ok {
		panic(fmt.Sprintf("invalid sslmode: %q. Valid options: disable, allow, prefer, require, verify-ca, verify-full", mode))
	}
}

func buildUrl(
	user string,
	password string,
	host string,
	port string,
	dbName string,
	sslMode SSLMode,
) string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   dbName,
	}

	validateSSLMode(sslMode)

	q := u.Query()
	q.Set("sslmode", string(sslMode))

	u.RawQuery = q.Encode()

	return u.String()
}

func NewPostgres(
	user string,
	password string,
	host string,
	port string,
	dbName string,
	sslMode SSLMode,
	maxOpenConns int,
	maxIdleConns int,
	connMaxLifetimeMins int,
) (*sql.DB, error) {
	dbUrl := buildUrl(user, password, host, port, dbName, sslMode)
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetimeMins) * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}
