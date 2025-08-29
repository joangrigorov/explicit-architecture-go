package cqrs

import (
	"app/internal/core/component/user/application/commands/complete_password_setup"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/idp"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	userEnt "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/events"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalCompletePasswordSetupCommand runs the CreateIdPUserCommandHandler in an Ent transaction,
// - an Ent transaction, handling commits and rollbacks.
// - uses TransactionalEventBus which flushes collected events only after successful command handling.
type TransactionalCompletePasswordSetupCommand struct {
	userRepository *usrAdapter.UserRepository
	eventBus       *events.SimpleEventBus
	idp            idp.IdentityProvider
	entClient      *userEnt.Client
}

func NewTransactionalCompletePasswordSetupCommand(
	userRepository *usrAdapter.UserRepository,
	eventBus *events.SimpleEventBus,
	idp idp.IdentityProvider,
	entClient *userEnt.Client,
) *TransactionalCompletePasswordSetupCommand {
	return &TransactionalCompletePasswordSetupCommand{
		userRepository: userRepository,
		eventBus:       eventBus,
		idp:            idp,
		entClient:      entClient,
	}
}

func (t *TransactionalCompletePasswordSetupCommand) Provide(ctx context.Context, command port.Command, next commands.Next) (err error) {
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

	handler := complete_password_setup.NewHandler(userRepository, t.idp, txEventBus)
	return handler.Handle(ctx, command.(complete_password_setup.Command))
}
