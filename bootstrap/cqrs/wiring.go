package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	. "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/idp"
	"app/internal/core/port/logging"
	cqrs2 "app/internal/infrastructure/component/user/cqrs"
	ent2 "app/internal/infrastructure/component/user/persistence/ent"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated"
	commandBus "app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/cqrs/commands/middleware"
	queryBus "app/internal/infrastructure/framework/cqrs/queries"
	middleware2 "app/internal/infrastructure/framework/cqrs/queries/middleware"
	"app/internal/infrastructure/framework/event_bus"

	"go.opentelemetry.io/otel/trace"
)

func WireCommands(
	logger logging.Logger,
	tracer trace.Tracer,
	bus *commandBus.SimpleCommandBus,
	userRepository *ent2.Repository,
	idp idp.IdentityProvider,
	eventBus *event_bus.SimpleEventBus,
	entClient *ent.Client,
) {
	bus.Use(middleware.Logger(logger))
	bus.Use(middleware.Tracing(tracer))

	commandBus.Register[RegisterUserCommand](bus, cqrs2.HandleRegisterUserCommand(userRepository, eventBus, entClient))
	commandBus.Register[ConfirmUserCommand](bus, cqrs2.TransactionalConfirmUserCommand(userRepository, idp, entClient))
	commandBus.Register[CreateIdPUserCommand](bus, cqrs2.HandleCreateIdPUserCommand(userRepository, idp, entClient))
}

func WireQueries(
	bus *queryBus.SimpleQueryBus,
	uq port.UserQueries,
	tracer trace.Tracer,
) {
	bus.Use(middleware2.Tracing(tracer))

	queryBus.Register[FindUserByIDQuery](
		bus,
		middleware2.ExecuteQuery[FindUserByIDQuery, *dto.UserDTO](NewFindUserByIDHandler(uq)),
	)
}
