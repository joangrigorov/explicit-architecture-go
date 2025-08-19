package fx

import (
	"app/bootstrap/cqrs"
	"app/bootstrap/events"
	cBus "app/internal/infrastructure/cqrs/commands"
	qBus "app/internal/infrastructure/cqrs/queries"
	"app/internal/infrastructure/event_bus"
	"app/internal/infrastructure/http"
	"app/internal/infrastructure/idp"
	"app/internal/infrastructure/logging/zap"
	"app/internal/infrastructure/observability/otel"
	"app/internal/infrastructure/persistence/ent/activity"
	"app/internal/infrastructure/persistence/ent/attendance"
	"app/internal/infrastructure/persistence/ent/user"
	"app/internal/infrastructure/uuid"
	"app/internal/infrastructure/validation"

	"go.uber.org/fx"
)

var cqrsAdapter = fx.Module("cqrs", fx.Provide(
	cBus.NewCommandBus,
	cBus.NewSimpleCommandBus,
	qBus.NewQueryBus,
	qBus.NewSimpleQueryBus,
), fx.Invoke(
	cqrs.WireCommands,
	cqrs.WireQueries,
))

var entActivityAdapter = fx.Module("persistence/ent/activity", fx.Provide(
	activity.NewConnection,
	activity.NewClient,
	activity.NewActivityRepository,
))

var entAttendanceAdapter = fx.Module("persistence/ent/attendance", fx.Provide(
	attendance.NewConnection,
	attendance.NewClient,
	attendance.NewAttendanceRepository,
))

var entUserAdapter = fx.Module("persistence/ent/user", fx.Provide(
	user.NewConnection,
	user.NewClient,
	user.NewRepository,
	user.NewConcreteRepository,
	user.NewQueries,
))

var persistenceAdapter = fx.Module("persistence",
	entActivityAdapter,
	entAttendanceAdapter,
	entUserAdapter,
)

var httpAdapter = fx.Module("http", fx.Provide(
	http.NewGinEngine,
	http.NewRouter,
))

var validationAdapter = fx.Module("validation", fx.Provide(
	validation.NewValidatorValidate,
	validation.NewTranslator,
), fx.Invoke(
	validation.RegisterRules,
))

var loggingAdapter = fx.Module("logging", fx.Provide(
	zap.NewZapLogger,
	zap.NewLogger,
), fx.Invoke(
	zap.ConfigureZap,
))

var idpAdapter = fx.Module("idp", fx.Provide(
	idp.NewGoCloakClient,
	idp.NewKeycloakIdentityProvider,
))

var eventBusAdapter = fx.Module("eventbus", fx.Provide(
	event_bus.NewEventBus,
	event_bus.NewSimpleEventBus,
), fx.Invoke(
	events.WireSubscribers,
))

var observabilityAdapter = fx.Module("observability", fx.Provide(
	otel.NewTracerProvider,
	otel.DefaultTracer,
), fx.Invoke(
	otel.RegisterTracer,
	otel.AddOpenTelemetryMiddleware,
))

var Infrastructure = fx.Module("infrastructure",
	cqrsAdapter,
	persistenceAdapter,
	httpAdapter,
	validationAdapter,
	loggingAdapter,
	idpAdapter,
	eventBusAdapter,
	observabilityAdapter,
	fx.Provide(uuid.NewGenerator),
)
