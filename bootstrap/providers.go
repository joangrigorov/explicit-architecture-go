package bootstrap

import (
	"app/config"
	"app/internal/infrastructure/observability/otel"
	"app/internal/infrastructure/persistence/ent/activity"
	"app/internal/infrastructure/persistence/ent/attendance"
	"app/internal/presentation/web/core/component/activity/v1/controllers/activities"
	"app/internal/presentation/web/infrastructure/framework/http"
	"app/internal/presentation/web/infrastructure/framework/validation"

	"go.uber.org/fx"
)

var providers = fx.Options(
	fx.Provide(
		// configuration providers
		config.NewConfig,

		// persistence providers
		activity.NewConnection,
		activity.NewClient,

		attendance.NewConnection,
		attendance.NewClient,

		// framework providers
		http.NewGinEngine,
		http.NewRouter,
		validation.NewValidatorValidate,
		validation.NewTranslator,

		// observability providers
		otel.NewTracerProvider,
		otel.DefaultTracer,

		// repository adapter providers
		activity.NewActivityRepository,
		attendance.NewAttendanceRepository,

		// web controller providers
		activities.NewController,
	),
)
