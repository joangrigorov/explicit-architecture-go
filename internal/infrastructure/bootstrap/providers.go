package bootstrap

import (
	"app/config"
	"app/internal/infrastructure/framework/http"
	"app/internal/infrastructure/framework/validation"
	"app/internal/infrastructure/persistence/ent"
	"app/internal/infrastructure/persistence/ent/blog"
	"app/internal/presentation/web/core/component/blog/v1/anonymous/controllers/post"
	"go.uber.org/fx"
)

var providers = fx.Options(
	fx.Provide(
		// configuration providers
		config.NewConfig,

		// database (ent) related providers
		ent.NewDB,
		blog.NewClient,

		// framework providers
		http.NewGinEngine,
		http.NewRouter,
		validation.NewValidatorValidate,
		validation.NewTranslator,

		// repository adapter providers
		blog.NewPostRepository,

		// web controller providers
		post.NewController,
	),
)
