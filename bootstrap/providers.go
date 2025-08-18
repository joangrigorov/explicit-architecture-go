package bootstrap

import (
	"app/config"
	cqrs "app/internal/core/component/user/application/queries"
	"app/internal/infrastructure/commands"
	"app/internal/infrastructure/events"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/idp"
	"app/internal/infrastructure/logging/zap"
	"app/internal/infrastructure/observability/otel"
	"app/internal/infrastructure/persistence/ent/activity"
	"app/internal/infrastructure/persistence/ent/attendance"
	"app/internal/infrastructure/persistence/ent/user"
	"app/internal/infrastructure/persistence/ent/user/queries"
	queryBus "app/internal/infrastructure/queries"
	"app/internal/presentation/api/core/component/activity/v1/controllers/activities"
	"app/internal/presentation/api/core/component/user/v1/controllers"
	"app/internal/presentation/api/infrastructure/framework/http"
	"app/internal/presentation/api/infrastructure/framework/validation"

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
		user.NewConnection,
		user.NewClient,

		// framework providers
		http.NewGinEngine,
		http.NewRouter,
		validation.NewValidatorValidate,
		validation.NewTranslator,
		uuid.NewGenerator,
		zap.NewZapLogger,
		zap.NewLogger,
		commands.NewCommandBus,
		commands.NewSimpleCommandBus,
		queryBus.NewSimpleQueryBus,
		queryBus.NewQueryBus,

		// keycloak providers
		idp.NewGoCloakClient,
		idp.NewKeycloakIdentityProvider,

		// event bus providers
		events.NewEventBus,
		events.NewSimpleEventBus,

		// observability providers
		otel.NewTracerProvider,
		otel.DefaultTracer,

		// repository adapter providers
		activity.NewActivityRepository,
		attendance.NewAttendanceRepository,
		user.NewRepository,
		user.NewConcreteRepository,

		// api controller providers
		activities.NewController,
		controllers.NewRegistrationController,

		// queries
		cqrs.NewFindUserByIDHandler,
		queries.NewUserQueries,
	),
)
