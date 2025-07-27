package blog

import (
	"app/internal/infrastructure/persistence/ent/generated/blog"
	"context"
	"go.uber.org/fx"
)

func MigrateSchema(lc fx.Lifecycle, client *blog.Client) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
	})
}
