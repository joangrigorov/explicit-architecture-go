package fx

import (
	"app/config"
	"app/internal/infrastructure/persistence/ent"
	"app/internal/infrastructure/persistence/ent/blog"
	"app/internal/presentation/web"
	"go.uber.org/fx"
)

var Providers = fx.Options(
	fx.Provide(
		web.NewRouter,
		blog.NewPostRepository,
		config.NewConfig,
		ent.NewDB,
		blog.NewClient,
	),
)
