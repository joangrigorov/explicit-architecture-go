package fx

import (
	cqrsBootstrap "app/bootstrap/cqrs"
	eventsWiring "app/bootstrap/events"
	activity "app/internal/infrastructure/component/activity/persistence/ent"
	attendance "app/internal/infrastructure/component/attendance/persistence/ent"
	userCQRS "app/internal/infrastructure/component/user/cqrs"
	user "app/internal/infrastructure/component/user/persistence/ent"
	cBus "app/internal/infrastructure/framework/cqrs/commands"
	qBus "app/internal/infrastructure/framework/cqrs/queries"
	"app/internal/infrastructure/framework/errors"
	"app/internal/infrastructure/framework/events"
	"app/internal/infrastructure/framework/http"
	"app/internal/infrastructure/framework/idp"
	"app/internal/infrastructure/framework/logging/zap"
	"app/internal/infrastructure/framework/mail"
	"app/internal/infrastructure/framework/observability/otel"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/framework/validation"

	"go.uber.org/fx"
)

var Http = fx.Module("http", fx.Provide(
	http.NewGinEngine,
	http.NewRouter,
))

var Logging = fx.Module("logging", fx.Provide(
	zap.NewZapLogger,
	zap.NewLogger,
), fx.Invoke(
	zap.ConfigureZap,
))

var Infrastructure = fx.Module("infrastructure",
	fx.Module("framework",
		Http,
		Logging,
		fx.Module("errors", fx.Provide(
			errors.NewErrorFactory,
		)),
		fx.Module("mail", fx.Provide(
			mail.NewDialer,
			mail.NewGomailMailer,
			mail.NewMailer,
			mail.NewTransactionalMailer,
		)),
		fx.Module("cqrs", fx.Provide(
			cBus.NewCommandBus,
			cBus.NewSimpleCommandBus,
			qBus.NewQueryBus,
			qBus.NewSimpleQueryBus,
		), fx.Invoke(
			cqrsBootstrap.WireCommands,
			cqrsBootstrap.WireQueries,
		)),
		fx.Module("validation", fx.Provide(
			validation.NewValidatorValidate,
			validation.NewTranslator,
		), fx.Invoke(
			validation.RegisterRules,
		)),
		fx.Module("idp", fx.Provide(
			idp.NewGoCloakClient,
			idp.NewKeycloakIdentityProvider,
		)),
		fx.Module("events", fx.Provide(
			events.NewEventBus,
			events.NewSimpleEventBus,
		), fx.Invoke(
			eventsWiring.WireSubscribers,
		)),
		fx.Module("observability", fx.Provide(
			otel.NewTracerProvider,
			otel.DefaultTracer,
		), fx.Invoke(
			otel.RegisterTracer,
			otel.AddOpenTelemetryMiddleware,
		)),
		fx.Module("uuid", fx.Provide(
			uuid.NewGenerator,
		)),
	),
	fx.Module("component",
		fx.Module("activity", fx.Module("persistence/ent", fx.Provide(
			activity.NewConnection,
			activity.NewClient,
			activity.NewActivityRepository,
		))),
		fx.Module("attendance", fx.Module("persistence/ent", fx.Provide(
			attendance.NewConnection,
			attendance.NewClient,
			attendance.NewAttendanceRepository,
		))),
		fx.Module("user", fx.Module("persistence/ent", fx.Provide(
			user.NewConnection,
			user.NewClient,
			user.NewRepository,
			user.NewConcreteRepository,
			user.NewQueries,
			user.NewConfirmationRepository,
			user.NewConcreteConfirmationRepository,
			userCQRS.NewTransactionalRegisterUserCommand,
			userCQRS.NewTransactionalInitiatePasswordSetupCommand,
			userCQRS.NewTransactionalCompletePasswordSetupCommand,
			userCQRS.NewTransactionalConfirmUserCommand,
		))),
	),
)
