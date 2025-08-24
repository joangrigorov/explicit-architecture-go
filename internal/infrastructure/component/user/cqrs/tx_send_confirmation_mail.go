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
	"context"
)

// TransactionalSendConfirmationMailCommand runs the SendConfirmationMailHandler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
// - uses TransactionalMailer which flushes collected mailables only after successful command handling.
func TransactionalSendConfirmationMailCommand(
	userRepository *usrAdapter.UserRepository,
	confirmRepository *usrAdapter.ConfirmationRepository,
	entClient *userEnt.Client,
	confirmMail mailables.ConfirmationMail,

	// framework
	uuidGenerator uuid.Generator,
	hmacGenerator hmac.Generator,
	mailer *mail.Mailer,
	logger logging.Logger,
) commands.Middleware {
	return func(ctx context.Context, command port.Command, next commands.Next) error {
		tx, err := entClient.Tx(ctx)
		if err != nil {
			return err
		}

		txMailer := mail.NewTransactionalMailer(mailer, logger)

		defer func() {
			if err == nil {
				txMailer.Flush()
			} else {
				txMailer.Reset()
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				_ = tx.Rollback()
				panic(r)
			} else if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()

		confirmSvc := services.NewConfirmationService(confirmRepository.WithTx(tx), uuidGenerator, hmacGenerator)
		mailerSvc := services.NewMailService(txMailer, confirmMail)

		handler := NewSendConfirmationMailHandler(userRepository.WithTx(tx), confirmSvc, mailerSvc)

		err = handler.Handle(ctx, command.(SendConfirmationMailCommand))
		return err
	}
}
