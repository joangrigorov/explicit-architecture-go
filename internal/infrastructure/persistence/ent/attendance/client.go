package attendance

import (
	"app/internal/infrastructure/persistence/ent/generated/attendances"

	"database/sql"

	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
)

func NewClient(db *sql.DB) *attendances.Client {
	return attendances.NewClient(attendances.Driver(entSql.OpenDB(dialect.Postgres, db)))
}
