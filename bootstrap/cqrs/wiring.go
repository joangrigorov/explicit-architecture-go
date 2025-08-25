package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	. "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/logging"
	"app/internal/infrastructure/component/user/cqrs"
	cBus "app/internal/infrastructure/framework/cqrs/commands"
	cMdwr "app/internal/infrastructure/framework/cqrs/commands/middleware"
	qBus "app/internal/infrastructure/framework/cqrs/queries"
	qMdwr "app/internal/infrastructure/framework/cqrs/queries/middleware"

	"go.opentelemetry.io/otel/trace"
)

func WireCommands(
	registerUserHandler *cqrs.TransactionalRegisterUserCommand,
	confirmUserHandler *cqrs.TransactionalConfirmUserCommand,
	createIdPUserHandler *cqrs.TransactionalCreateIdPUserCommand,
	sndCnfrmMailHandler *cqrs.TransactionalSendConfirmationMailCommand,

	// framework
	logger logging.Logger,
	tracer trace.Tracer,
	bus *cBus.SimpleCommandBus,
) {
	bus.Use(cMdwr.Logger(logger))
	bus.Use(cMdwr.Tracing(tracer))

	cBus.Register[RegisterUserCommand](bus, registerUserHandler.Provide)
	cBus.Register[ConfirmUserCommand](bus, confirmUserHandler.Provide)
	cBus.Register[CreateIdPUserCommand](bus, createIdPUserHandler.Provide)
	cBus.Register[SendConfirmationMailCommand](bus, sndCnfrmMailHandler.Provide)
}

func WireQueries(
	bus *qBus.SimpleQueryBus,
	uq port.UserQueries,
	tracer trace.Tracer,
) {
	bus.Use(qMdwr.Tracing(tracer))

	qBus.Register[FindUserByIDQuery](
		bus,
		qMdwr.ExecuteQuery[FindUserByIDQuery, *dto.UserDTO](NewFindUserByIDHandler(uq)),
	)
}
