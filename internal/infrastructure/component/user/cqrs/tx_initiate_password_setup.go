package cqrs

import (
	"app/internal/core/component/user/application/commands/initiate_password_setup"
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/application/services"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	"app/internal/core/port/hmac"
	"app/internal/core/port/logging"
	"app/internal/core/port/uuid"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/events"
	"app/internal/infrastructure/framework/mail"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalInitiatePasswordSetupCommand runs the Handler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
// - uses TransactionalMailer which flushes collected mailables only after successful command handling.
type TransactionalInitiatePasswordSetupCommand struct {
	userRepository    *usrAdapter.UserRepository
	confirmRepository *usrAdapter.ConfirmationRepository
	eventBus          *events.SimpleEventBus
	entClient         *userEnt.Client
	confirmMail       mailables.PasswordSetupMail
	uuidGenerator     uuid.Generator
	hmacGenerator     hmac.Generator
	mailer            *mail.Mailer
	logger            logging.Logger
	errors            errors.ErrorFactory
}

func NewTransactionalInitiatePasswordSetupCommand(
	userRepository *usrAdapter.UserRepository,
	confirmRepository *usrAdapter.ConfirmationRepository,
	eventBus *events.SimpleEventBus,
	entClient *userEnt.Client,
	confirmMail mailables.PasswordSetupMail,
	uuidGenerator uuid.Generator,
	hmacGenerator hmac.Generator,
	mailer *mail.Mailer,
	logger logging.Logger,
	errors errors.ErrorFactory,
) *TransactionalInitiatePasswordSetupCommand {
	return &TransactionalInitiatePasswordSetupCommand{
		userRepository:    userRepository,
		confirmRepository: confirmRepository,
		eventBus:          eventBus,
		entClient:         entClient,
		confirmMail:       confirmMail,
		uuidGenerator:     uuidGenerator,
		hmacGenerator:     hmacGenerator,
		mailer:            mailer,
		logger:            logger,
		errors:            errors,
	}
}

func (t *TransactionalInitiatePasswordSetupCommand) Provide(ctx context.Context, command port.Command, _ commands.Next) (err error) {
	txEventBus := events.NewTransactionalEventBus(t.eventBus)
	defer events.CloseEventBus(ctx, txEventBus, &err)

	mailer := mail.NewTransactionalMailer(t.mailer, t.logger)
	defer mail.CloseMailer(mailer, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	confirmationRepository := t.confirmRepository.
		WithTx(tx).
		WithEventBus(txEventBus)

	confirmSvc := services.NewConfirmationService(confirmationRepository, t.uuidGenerator, t.hmacGenerator)
	mailerSvc := services.NewMailService(mailer, t.confirmMail)

	userRepository := t.userRepository.
		WithTx(tx).
		WithEventBus(txEventBus)

	handler := initiate_password_setup.NewHandler(userRepository, confirmSvc, mailerSvc, t.errors)
	return handler.Handle(ctx, command.(initiate_password_setup.Command))
}
