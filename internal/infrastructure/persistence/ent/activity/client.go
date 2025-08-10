package activity

import (
	"app/config"
	"app/internal/infrastructure/persistence/ent/generated/activity"

	"database/sql"

	entSql "entgo.io/ent/dialect/sql"
)

func NewClient(db *sql.DB, cfg *config.Config) *activity.Client {
	return activity.NewClient(activity.Driver(entSql.OpenDB(cfg.DB.Driver, db)))
}
