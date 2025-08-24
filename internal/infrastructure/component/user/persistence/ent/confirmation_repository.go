package ent

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/uuid"
	"context"
)

type ConfirmationRepository struct {
	entTx     *generated.Tx
	entClient *generated.Client
}

func NewConfirmationRepository(r *ConfirmationRepository) repositories.ConfirmationRepository {
	return r
}

func NewConcreteConfirmationRepository(entClient *generated.Client) *ConfirmationRepository {
	return &ConfirmationRepository{entClient: entClient}
}

func (r *ConfirmationRepository) Create(ctx context.Context, c *domain.Confirmation) error {
	_, err := r.client().
		Create().
		SetID(uuid.Parse(c.ID.String())).
		SetUserID(uuid.Parse(c.UserID.String())).
		SetHmacSecret(c.HMACSecret).
		SetCreatedAt(c.CreatedAt).
		Save(ctx)

	return err
}

func (r *ConfirmationRepository) GetByID(ctx context.Context, id domain.ConfirmationID) (*domain.Confirmation, error) {
	dto, err := r.client().Get(ctx, uuid.Parse(id.String()))

	if err != nil {
		return nil, err
	}

	return mapConfirmationAggregate(dto), nil
}

func (r *ConfirmationRepository) Expire(ctx context.Context, c *domain.Confirmation) error {
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
