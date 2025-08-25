package cqrs

import (
	. "app/internal/core/component/user/application/commands"
	port "app/internal/core/port/cqrs"
	"app/internal/core/port/idp"
	usrAdapter "app/internal/infrastructure/component/user/persistence/ent"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/cqrs/commands"
	"app/internal/infrastructure/framework/persistence/pgsql"
	"context"
)

// TransactionalConfirmUserCommand runs the ConfirmUserCommandHandler in an Ent transaction,
// and it handles commit and rollbacks.
type TransactionalConfirmUserCommand struct {
	userRepository *usrAdapter.UserRepository
	entClient      *ent.Client
	idp            idp.IdentityProvider
}

func NewTransactionalConfirmUserCommand(
	userRepository *usrAdapter.UserRepository,
	entClient *ent.Client,
	idp idp.IdentityProvider,
) *TransactionalConfirmUserCommand {
	return &TransactionalConfirmUserCommand{userRepository: userRepository, entClient: entClient, idp: idp}
}

func (t *TransactionalConfirmUserCommand) Provide(ctx context.Context, command port.Command, _ commands.Next) (err error) {
	tx, err := t.entClient.Tx(ctx)
	if err != nil {
		return err
	}
	defer pgsql.CloseTx(tx, &err)

	handler := NewConfirmUserCommandHandler(t.userRepository.WithTx(tx), t.idp)
	return handler.Handle(ctx, command.(ConfirmUserCommand))
}
