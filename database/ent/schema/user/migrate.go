package user

import (
	"app/internal/infrastructure/persistence/ent/generated/user"
	"context"

	"go.uber.org/fx"
)

func MigrateSchema(lc fx.Lifecycle, client *user.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
	})
}
