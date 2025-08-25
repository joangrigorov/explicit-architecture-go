package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/application/services"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/hmac"
	"app/internal/core/port/logging"
	"app/internal/core/port/uuid"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/mail"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalSendConfirmationMailCommand runs the SendConfirmationMailHandler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
// - uses TransactionalMailer which flushes collected mailables only after successful command handling.
type TransactionalSendConfirmationMailCommand struct {
	userRepository    *usrAdapter.UserRepository
	confirmRepository *usrAdapter.ConfirmationRepository
	entClient         *userEnt.Client
	confirmMail       mailables.ConfirmationMail
	uuidGenerator     uuid.Generator
	hmacGenerator     hmac.Generator
	mailer            *mail.Mailer
	logger            logging.Logger
}

func NewTransactionalSendConfirmationMailCommand(
	userRepository *usrAdapter.UserRepository,
	confirmRepository *usrAdapter.ConfirmationRepository,
	entClient *userEnt.Client,
	confirmMail mailables.ConfirmationMail,
	uuidGenerator uuid.Generator,
	hmacGenerator hmac.Generator,
	mailer *mail.Mailer,
	logger logging.Logger,
) *TransactionalSendConfirmationMailCommand {
	return &TransactionalSendConfirmationMailCommand{
		userRepository:    userRepository,
		confirmRepository: confirmRepository,
		entClient:         entClient,
		confirmMail:       confirmMail,
		uuidGenerator:     uuidGenerator,
		hmacGenerator:     hmacGenerator,
		mailer:            mailer,
		logger:            logger,
	}
}

func (t *TransactionalSendConfirmationMailCommand) Provide(ctx context.Context, command port.Command, _ commands.Next) (err error) {
	mailer := mail.NewTransactionalMailer(t.mailer, t.logger)
	defer mail.CloseMailer(mailer, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	confirmSvc := services.NewConfirmationService(t.confirmRepository.WithTx(tx), t.uuidGenerator, t.hmacGenerator)
	mailerSvc := services.NewMailService(mailer, t.confirmMail)

	handler := NewSendConfirmationMailHandler(t.userRepository.WithTx(tx), confirmSvc, mailerSvc)
	return handler.Handle(ctx, command.(SendConfirmationMailCommand))
}
