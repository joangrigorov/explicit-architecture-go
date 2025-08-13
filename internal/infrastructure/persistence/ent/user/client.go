package user

import (
	"app/config"
	"app/internal/infrastructure/persistence/ent/generated/user"
	"app/internal/infrastructure/persistence/pgsql"
	"database/sql"
	"fmt"

	entSql "entgo.io/ent/dialect/sql"
)

type Connection *sql.DB

func NewConnection(cfg *config.Config) (Connection, error) {
	cfgDb := cfg.DB.User

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

func NewClient(db Connection, cfg *config.Config) *user.Client {
	return user.NewClient(
		user.Driver(
			entSql.OpenDB(cfg.DB.User.Driver, db),
		),
	)
}
