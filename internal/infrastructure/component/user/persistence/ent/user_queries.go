package ent

import (
	"app/internal/core/component/user/application/queries/find_user_by_id"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"app/internal/infrastructure/framework/uuid"
	"context"
)

type Queries struct {
	client *generated.Client
}

func NewQueries(client *generated.Client) find_user_by_id.UserQueries {
	return &Queries{client: client}
}

func (u *Queries) FindById(ctx context.Context, id string) (*find_user_by_id.UserDTO, error) {
	entDto, err := u.client.User.
		Query().
		Where(ent.ID(uuid.Parse(id)), ent.DeletedAtIsNil()).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return &find_user_by_id.UserDTO{
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
