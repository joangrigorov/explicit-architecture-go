package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	"app/internal/core/component/user/application/mailables"
	. "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/hmac"
	"app/internal/core/port/idp"
	"app/internal/core/port/logging"
	"app/internal/core/port/uuid"
	"app/internal/infrastructure/component/user/cqrs"
	user "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	commandBus "app/internal/infrastructure/framework/cqrs/commands"
	cMdwr "app/internal/infrastructure/framework/cqrs/commands/middleware"
	queryBus "app/internal/infrastructure/framework/cqrs/queries"
	qMdwr "app/internal/infrastructure/framework/cqrs/queries/middleware"
	"app/internal/infrastructure/framework/event_bus"
	"app/internal/infrastructure/framework/mail"

	"go.opentelemetry.io/otel/trace"
)

func WireCommands(
	userRepository *user.UserRepository,
	confirmRepository *user.ConfirmationRepository,
	confirmMail mailables.ConfirmationMail,
	userEntClient *userEnt.Client,

	// framework
	logger logging.Logger,
	tracer trace.Tracer,
	bus *commandBus.SimpleCommandBus,
	idp idp.IdentityProvider,
	eventBus *event_bus.SimpleEventBus,
	uuidGenerator uuid.Generator,
	hmacGenerator hmac.Generator,
	mailer *mail.Mailer,
) {
	bus.Use(cMdwr.Logger(logger))
	bus.Use(cMdwr.Tracing(tracer))

	commandBus.Register[RegisterUserCommand](bus, cqrs.TransactionalRegisterUserCommand(userRepository, eventBus, userEntClient))
	commandBus.Register[ConfirmUserCommand](bus, cqrs.TransactionalConfirmUserCommand(userRepository, idp, userEntClient))
	commandBus.Register[CreateIdPUserCommand](bus, cqrs.TransactionalCreateIdPUserCommand(userRepository, eventBus, idp, userEntClient))
	commandBus.Register[SendConfirmationMailCommand](bus, cqrs.TransactionalSendConfirmationMailCommand(
		userRepository,
		confirmRepository,
		userEntClient,
		confirmMail,
		uuidGenerator,
		hmacGenerator,
		mailer,
		logger,
	))
}

func WireQueries(
	bus *queryBus.SimpleQueryBus,
	uq port.UserQueries,
	tracer trace.Tracer,
) {
	bus.Use(qMdwr.Tracing(tracer))

	queryBus.Register[FindUserByIDQuery](
		bus,
		qMdwr.ExecuteQuery[FindUserByIDQuery, *dto.UserDTO](NewFindUserByIDHandler(uq)),
	)
}
