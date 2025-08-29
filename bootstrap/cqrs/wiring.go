package cqrs

import (
	"app/internal/core/component/user/application/commands/complete_password_setup"
	"app/internal/core/component/user/application/commands/confirm_user"
	"app/internal/core/component/user/application/commands/initiate_password_setup"
	"app/internal/core/component/user/application/commands/register_user"
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
	registerUser *cqrs.TransactionalRegisterUserCommand,
	initPsswdSetup *cqrs.TransactionalInitiatePasswordSetupCommand,
	completePsswdSetup *cqrs.TransactionalCompletePasswordSetupCommand,
	cfrmUser *cqrs.TransactionalConfirmUserCommand,

	// framework
	logger logging.Logger,
	tracer trace.Tracer,
	bus *cBus.SimpleCommandBus,
) {
	bus.Use(cMdwr.Logger(logger))
	bus.Use(cMdwr.Tracing(tracer))

	cBus.Register[register_user.Command](bus, registerUser.Provide)
	cBus.Register[initiate_password_setup.Command](bus, initPsswdSetup.Provide)
	cBus.Register[complete_password_setup.Command](bus, completePsswdSetup.Provide)
	cBus.Register[confirm_user.Command](bus, cfrmUser.Provide)
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
