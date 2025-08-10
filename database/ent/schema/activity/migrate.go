package activity

import (
	"app/internal/infrastructure/persistence/ent/generated/activity"
	"context"

	"go.uber.org/fx"
)

func MigrateSchema(lc fx.Lifecycle, client *activity.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
	})
}
