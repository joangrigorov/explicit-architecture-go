package bootstrap

import (
	"app/database/ent/schema/activity"
	"app/database/ent/schema/attendance"
	"app/database/ent/schema/user"
	"app/internal/infrastructure/events/subscribers/create_keycloak_user"
	"app/internal/infrastructure/logging/zap"
	"app/internal/infrastructure/observability/otel"
	"app/internal/presentation/api/core"
	"app/internal/presentation/api/infrastructure/framework/validation"

	"go.uber.org/fx"
)

var bootstraps = fx.Options(
	fx.Invoke(
		// observability
		otel.RegisterTracer,
		otel.AddOpenTelemetryMiddleware,
		zap.ConfigureZap,

		// Migrations
		activity.MigrateSchema,
		attendance.MigrateSchema,
		user.MigrateSchema,

		// bootstrap
		core.RegisterRoutes,
		validation.RegisterRules,

		// event subscribers
		create_keycloak_user.Register,

		// initiate
		runServer,
	),
)
