package bootstrap

import (
	"app/database/ent/schema/activity"
	"app/database/ent/schema/attendance"
	"app/database/ent/schema/user"
	"app/internal/infrastructure/observability/otel"
	"app/internal/presentation/web/core"
	"app/internal/presentation/web/infrastructure/framework/validation"

	"go.uber.org/fx"
)

var bootstraps = fx.Options(
	fx.Invoke(
		// observability
		otel.RegisterTracer,
		otel.AddOpenTelemetryMiddleware,

		// Migrations
		activity.MigrateSchema,
		attendance.MigrateSchema,
		user.MigrateSchema,

		// bootstrap
		core.RegisterRoutes,
		validation.RegisterRules,

		// initiate
		runServer,
	),
)
