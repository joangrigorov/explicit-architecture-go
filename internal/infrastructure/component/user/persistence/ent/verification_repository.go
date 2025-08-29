package ent

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/verification"
	"app/internal/core/port/events"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/framework/uuid"
	"context"
)

type VerificationRepository struct {
	eventBus  events.EventBus
	entTx     *generated.Tx
	entClient *generated.Client
}

func (r *VerificationRepository) Create(ctx context.Context, c *verification.Verification) error {
	_, err := r.client().
		Create().
		SetID(uuid.Parse(c.ID.String())).
		SetUserID(uuid.Parse(c.UserID.String())).
		SetCsrfToken(c.CSRFToken.Encode()).
		SetExpiresAt(c.ExpiresAt).
		SetCreatedAt(c.CreatedAt).
		Save(ctx)

	// TODO offload events into event bus

	return err
}

func (r *VerificationRepository) GetByID(ctx context.Context, id verification.ID) (*verification.Verification, error) {
	dto, err := r.client().Get(ctx, uuid.Parse(id.String()))

	if err != nil {
		return nil, err
	}

	return mapVerificationAggregate(dto)
}

func (r *VerificationRepository) Expire(ctx context.Context, c *verification.Verification) error {
	return r.client().DeleteOneID(uuid.Parse(c.ID.String())).Exec(ctx)
}

func (r *VerificationRepository) client() *generated.VerificationClient {
	if r.entTx != nil {
		return r.entTx.Verification
	}

	if r.entClient != nil {
		return r.entClient.Verification
	}

	panic("ent client not initialized")
}

func (r *VerificationRepository) WithTx(tx *generated.Tx) *VerificationRepository {
	return &VerificationRepository{entTx: tx, entClient: r.entClient}
}

func (r *VerificationRepository) WithEventBus(eventBus events.EventBus) *VerificationRepository {
	return &VerificationRepository{entTx: r.entTx, entClient: r.entClient, eventBus: eventBus}
}

func NewConfirmationRepository(r *VerificationRepository) repositories.VerificationRepository {
	return r
}

func NewConcreteConfirmationRepository(entClient *generated.Client, eventBus events.EventBus) *VerificationRepository {
	return &VerificationRepository{entClient: entClient, eventBus: eventBus}
}
