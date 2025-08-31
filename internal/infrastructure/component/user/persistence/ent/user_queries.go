package ent

import (
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"app/internal/infrastructure/framework/uuid"
	"context"
)

type UserQueries struct {
	client *generated.Client
}

func NewUserQueries(client *generated.Client) port.UserQueries {
	return &UserQueries{client: client}
}

func (u *UserQueries) FindByID(ctx context.Context, id string) (*dto.UserDTO, error) {
	entDto, err := u.client.User.
		Query().
		Where(ent.ID(uuid.Parse(id)), ent.DeletedAtIsNil()).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:          entDto.ID.String(),
		Email:       entDto.Email,
		Username:    entDto.Username,
		FirstName:   entDto.FirstName,
		LastName:    entDto.LastName,
		Role:        entDto.Role.String(),
		ConfirmedAt: entDto.ConfirmedAt,
		IdPUserID:   entDto.IdpUserID,
		CreatedAt:   entDto.CreatedAt,
		UpdatedAt:   entDto.UpdatedAt,
	}, nil
}
