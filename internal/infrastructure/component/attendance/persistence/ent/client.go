package ent

import (
	"app/config"
	"app/internal/infrastructure/component/attendance/persistence/ent/generated"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"fmt"

	"database/sql"

	entSql "entgo.io/ent/dialect/sql"
)

type Connection *sql.DB

func NewConnection(cfg *config.Config) (Connection, error) {
	cfgDb := cfg.DB.Attendance

	if cfgDb.Driver == "postgres" {
		return pgsql.NewPostgres(
			cfgDb.User,
			cfgDb.Password,
			cfgDb.Host,
			cfgDb.Port,
			cfgDb.Database,
			pgsql.SSLMode(cfgDb.SslMode),
			cfgDb.MaxOpenConns,
			cfgDb.MaxIdleConns,
			cfgDb.ConnMaxLifetime,
		)
	}

	panic(fmt.Sprintf("unsupported driver %s", cfgDb.Driver))
}

func NewClient(db Connection, cfg *config.Config) *generated.Client {
	return generated.NewClient(
		generated.Driver(
			entSql.OpenDB(cfg.DB.Attendance.Driver, db),
		),
	)
}
