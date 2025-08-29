package ent

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/confirmation"
	"app/internal/core/port/events"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/uuid"
	"context"
)

type ConfirmationRepository struct {
	eventBus  events.EventBus
	entTx     *generated.Tx
	entClient *generated.Client
}

func (r *ConfirmationRepository) Create(ctx context.Context, c *confirmation.Confirmation) error {
	_, err := r.client().
		Create().
		SetID(uuid.Parse(c.ID.String())).
		SetUserID(uuid.Parse(c.UserID.String())).
		SetHmacSecret(c.HMACSecret).
		SetCreatedAt(c.CreatedAt).
		Save(ctx)

	return err
}

func (r *ConfirmationRepository) GetByID(ctx context.Context, id confirmation.ID) (*confirmation.Confirmation, error) {
	dto, err := r.client().Get(ctx, uuid.Parse(id.String()))

	if err != nil {
		return nil, err
	}

	return mapConfirmationAggregate(dto), nil
}

func (r *ConfirmationRepository) Expire(ctx context.Context, c *confirmation.Confirmation) error {
	return r.client().DeleteOneID(uuid.Parse(c.ID.String())).Exec(ctx)
}

func (r *ConfirmationRepository) client() *generated.ConfirmationClient {
	if r.entTx != nil {
		return r.entTx.Confirmation
	}

	if r.entClient != nil {
		return r.entClient.Confirmation
	}

	panic("ent client not initialized")
}

func (r *ConfirmationRepository) WithTx(tx *generated.Tx) *ConfirmationRepository {
	return &ConfirmationRepository{entTx: tx, entClient: r.entClient}
}

func (r *ConfirmationRepository) WithEventBus(eventBus events.EventBus) *ConfirmationRepository {
	return &ConfirmationRepository{entTx: r.entTx, entClient: r.entClient, eventBus: eventBus}
}

func NewConfirmationRepository(r *ConfirmationRepository) repositories.ConfirmationRepository {
	return r
}

func NewConcreteConfirmationRepository(entClient *generated.Client, eventBus events.EventBus) *ConfirmationRepository {
	return &ConfirmationRepository{entClient: entClient, eventBus: eventBus}
}
