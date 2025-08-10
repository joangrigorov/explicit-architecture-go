package activity

import (
	"app/internal/infrastructure/persistence/ent/generated/activities"

	"database/sql"

	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
)

func NewClient(db *sql.DB) *activities.Client {
	return activities.NewClient(activities.Driver(entSql.OpenDB(dialect.Postgres, db)))
}
