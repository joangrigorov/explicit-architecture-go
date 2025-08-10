package attendance

import (
	"app/internal/infrastructure/persistence/ent/generated/attendance"

	"database/sql"

	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
)

func NewClient(db *sql.DB) *attendance.Client {
	return attendance.NewClient(attendance.Driver(entSql.OpenDB(dialect.Postgres, db)))
}
