package cqrs

import (
	"app/bootstrap/cqrs/decorators"
	. "app/internal/core/component/user/application/commands"
	. "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/idp"
	"app/internal/core/port/logging"
	commandBus "app/internal/infrastructure/cqrs/commands"
	cm "app/internal/infrastructure/cqrs/commands/middleware"
	queryBus "app/internal/infrastructure/cqrs/queries"
	qm "app/internal/infrastructure/cqrs/queries/middleware"
	"app/internal/infrastructure/event_bus"
	ent "app/internal/infrastructure/persistence/ent/generated/user"
	"app/internal/infrastructure/persistence/ent/user"

	"go.opentelemetry.io/otel/trace"
)

func WireCommands(
	logger logging.Logger,
	tracer trace.Tracer,
	bus *commandBus.SimpleCommandBus,
	userRepository *user.Repository,
	idp idp.IdentityProvider,
	eventBus *event_bus.SimpleEventBus,
	entClient *ent.Client,
) {
	bus.Use(cm.Logger(logger))
	bus.Use(cm.Tracing(tracer))

	commandBus.Register[RegisterUserCommand](bus, decorators.HandleRegisterUserCommand(userRepository, eventBus, entClient))
	commandBus.Register[ConfirmUserCommand](bus, decorators.TransactionalConfirmUserCommand(userRepository, idp, entClient))
	commandBus.Register[CreateIdPUserCommand](bus, decorators.HandleCreateIdPUserCommand(userRepository, idp, entClient))
}

func WireQueries(
	bus *queryBus.SimpleQueryBus,
	uq port.UserQueries,
	tracer trace.Tracer,
) {
	bus.Use(qm.Tracing(tracer))

	queryBus.Register[FindUserByIDQuery](
		bus,
		qm.ExecuteQuery[FindUserByIDQuery, *port.UserDTO](NewFindUserByIDHandler(uq)),
	)
}
