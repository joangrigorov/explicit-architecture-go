package ent

import (
	"app/config"

	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	if cfg.DB.Driver == "postgres" {
		return newPG(cfg)
	}

	panic(fmt.Sprintf("unsupported driver %s", cfg.DB.Driver))
}

func newPG(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DB.Driver, cfg.PG.URL)

	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}
