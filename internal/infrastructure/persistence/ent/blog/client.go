package blog

import (
	"app/internal/infrastructure/persistence/ent/generated/blog"

	"database/sql"
	"entgo.io/ent/dialect"
	entSql "entgo.io/ent/dialect/sql"
)

func NewClient(db *sql.DB) *blog.Client {
	return blog.NewClient(blog.Driver(entSql.OpenDB(dialect.Postgres, db)))
}
