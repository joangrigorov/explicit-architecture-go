package ent

import (
	"app/internal/core/component/user/application/queries/port"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	"app/internal/infrastructure/component/user/persistence/ent/generated/verification"
	"app/internal/infrastructure/framework/uuid"
	"context"
)

type VerificationQueries struct {
	client *generated.Client
}

func (u *VerificationQueries) FindByID(ctx context.Context, id string) (*port.VerificationDTO, error) {
	entDto, err := u.client.Verification.Get(ctx, uuid.Parse(id))
	if err != nil {
		return nil, err
	}

	return mapFromEnt(entDto), nil
}

func (u *VerificationQueries) FindByUserID(ctx context.Context, userID string) (*port.VerificationDTO, error) {
	entDto, err := u.client.Verification.
		Query().
		Where(verification.UserID(uuid.Parse(userID))).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapFromEnt(entDto), nil
}

func mapFromEnt(entDto *generated.Verification) *port.VerificationDTO {
	return &port.VerificationDTO{
		ID:              entDto.ID.String(),
		UserID:          entDto.UserID.String(),
		UserEmailMasked: entDto.UserEmailMasked,
		CSRFToken:       entDto.CsrfToken,
		ExpiresAt:       entDto.ExpiresAt,
		UsedAt:          entDto.UsedAt,
		CreatedAt:       entDto.CreatedAt,
	}
}

func NewVerificationQueries(client *generated.Client) port.VerificationQueries {
	return &VerificationQueries{client: client}
}
