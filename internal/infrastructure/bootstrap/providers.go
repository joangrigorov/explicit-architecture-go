package bootstrap

import (
	"app/config"
	"app/internal/infrastructure/framework/http"
	"app/internal/infrastructure/persistence/ent"
	"app/internal/infrastructure/persistence/ent/blog"
	"go.uber.org/fx"
)

var providers = fx.Options(
	fx.Provide(
		blog.NewPostRepository,
		config.NewConfig,
		ent.NewDB,
		blog.NewClient,
		http.NewGinEngine,
		http.NewRouter,
	),
)
