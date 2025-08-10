package attendance

import (
	"app/internal/infrastructure/persistence/ent/generated/attendances"
	"context"

	"go.uber.org/fx"
)

func MigrateSchema(lc fx.Lifecycle, client *attendances.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
	})
}
