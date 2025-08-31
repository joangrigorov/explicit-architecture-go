package cqrs

import (
	"app/internal/core/component/user/application/commands/complete_password_setup"
	"app/internal/core/component/user/application/commands/confirm_user"
	"app/internal/core/component/user/application/commands/initiate_password_setup"
	"app/internal/core/component/user/application/commands/register_user"
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/find_user_by_id"
	"app/internal/core/component/user/application/queries/get_verification_preflight"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/errors"
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
	vq port.VerificationQueries,
	tracer trace.Tracer,
	errors errors.ErrorFactory,
) {
	bus.Use(qMdwr.Tracing(tracer))

	qBus.Register[find_user_by_id.Query](
		bus,
		qMdwr.ExecuteQuery[find_user_by_id.Query, *dto.UserDTO](find_user_by_id.NewHandler(uq)),
	)

	qBus.Register[get_verification_preflight.Query](
		bus,
		qMdwr.ExecuteQuery[get_verification_preflight.Query, *dto.PreflightDTO](get_verification_preflight.NewHandler(vq, errors)),
	)
}
