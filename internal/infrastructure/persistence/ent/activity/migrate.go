package activity

import (
	"app/internal/infrastructure/persistence/ent/generated/activities"
	"context"

	"go.uber.org/fx"
)

func MigrateSchema(lc fx.Lifecycle, client *activities.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
	})
}
