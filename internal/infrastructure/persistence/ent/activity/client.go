package activity

import (
	"app/config"
	"app/internal/infrastructure/persistence/ent/generated/activity"
	"app/internal/infrastructure/persistence/pgsql"
	"database/sql"
	"fmt"

	entSql "entgo.io/ent/dialect/sql"
)

type Connection *sql.DB

func NewConnection(cfg *config.Config) (Connection, error) {
	cfgDb := cfg.DB.Activity
	
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

func NewClient(db Connection, cfg *config.Config) *activity.Client {
	return activity.NewClient(
		activity.Driver(
			entSql.OpenDB(cfg.DB.Activity.Driver, db),
		),
	)
}
