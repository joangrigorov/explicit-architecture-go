package cqrs

import (
	"app/config/api"
	"app/internal/core/component/user/application/commands/confirm_user"
	"app/internal/core/component/user/application/mailables"
	"app/internal/core/component/user/application/services"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	"app/internal/core/port/idp"
	"app/internal/core/port/logging"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/events"
	"app/internal/infrastructure/framework/mail"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalConfirmUserCommand runs the Handler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
type TransactionalConfirmUserCommand struct {
	userRepository    *usrAdapter.UserRepository
	passwordSetupMail mailables.PasswordSetupMail
	userConfirmedMail mailables.UserConfirmedMail
	eventBus          *events.SimpleEventBus
	mailer            *mail.Mailer
	logger            logging.Logger
	idp               idp.IdentityProvider
	entClient         *userEnt.Client
	errors            errors.ErrorFactory
	cfg               api.Config
}

func NewTransactionalConfirmUserCommand(
	userRepository *usrAdapter.UserRepository,
	passwordSetupMail mailables.PasswordSetupMail,
	userConfirmedMail mailables.UserConfirmedMail,
	eventBus *events.SimpleEventBus,
	mailer *mail.Mailer,
	logger logging.Logger,
	identityProvider idp.IdentityProvider,
	entClient *userEnt.Client,
	errors errors.ErrorFactory,
	cfg api.Config,
) *TransactionalConfirmUserCommand {
	return &TransactionalConfirmUserCommand{
		userRepository:    userRepository,
		passwordSetupMail: passwordSetupMail,
		userConfirmedMail: userConfirmedMail,
		eventBus:          eventBus,
		mailer:            mailer,
		logger:            logger,
		entClient:         entClient,
		errors:            errors,
		idp:               identityProvider,
		cfg:               cfg,
	}
}

func (t *TransactionalConfirmUserCommand) Provide(ctx context.Context, command port.Command, _ commands.Next) (err error) {
	txEventBus := events.NewTransactionalEventBus(t.eventBus)
	defer events.CloseEventBus(ctx, txEventBus, &err)

	mailer := mail.NewTransactionalMailer(t.mailer, t.logger)
	defer mail.CloseMailer(mailer, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	mailerSvc := services.NewMailService(mailer, t.passwordSetupMail, t.userConfirmedMail, t.errors)

	userRepository := t.userRepository.
		WithTx(tx).
		WithEventBus(txEventBus)

	handler := confirm_user.NewHandler(userRepository, t.idp, mailerSvc, t.errors, t.cfg)
	return handler.Handle(ctx, command.(confirm_user.Command))
}
