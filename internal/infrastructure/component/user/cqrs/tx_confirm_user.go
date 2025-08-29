package cqrs

import (
	"app/internal/core/component/user/application/commands/confirm_user"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/errors"
	"app/internal/core/port/idp"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/events"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalConfirmUserCommand runs the Handler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
type TransactionalConfirmUserCommand struct {
	userRepository *usrAdapter.UserRepository
	eventBus       *events.SimpleEventBus
	idp            idp.IdentityProvider
	entClient      *userEnt.Client
	errors         errors.ErrorFactory
}

func NewTransactionalConfirmUserCommand(
	userRepository *usrAdapter.UserRepository,
	eventBus *events.SimpleEventBus,
	identityProvider idp.IdentityProvider,
	entClient *userEnt.Client,
	errors errors.ErrorFactory,
) *TransactionalConfirmUserCommand {
	return &TransactionalConfirmUserCommand{
		userRepository: userRepository,
		eventBus:       eventBus,
		entClient:      entClient,
		errors:         errors,
		idp:            identityProvider,
	}
}

func (t *TransactionalConfirmUserCommand) Provide(ctx context.Context, command port.Command, _ commands.Next) (err error) {
	txEventBus := events.NewTransactionalEventBus(t.eventBus)
	defer events.CloseEventBus(ctx, txEventBus, &err)

	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	userRepository := t.userRepository.
		WithTx(tx).
		WithEventBus(txEventBus)

	handler := confirm_user.NewHandler(userRepository, t.idp, t.errors)
	return handler.Handle(ctx, command.(confirm_user.Command))
}
