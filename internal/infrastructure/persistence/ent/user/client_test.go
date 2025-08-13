package user

import (
	"app/config"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestNewClient(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	cfg := &config.Config{
		DB: config.DB{
			User: config.DBUser{
				Driver: "sqlite3",
			},
		},
	}

	client := NewClient(db, cfg)
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}
