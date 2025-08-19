package user

import (
	"app/internal/core/component/user/application/queries/port"
	"app/internal/infrastructure/persistence/ent/generated/user"
	ent "app/internal/infrastructure/persistence/ent/generated/user/user"
	"app/internal/infrastructure/uuid"
	"context"
)

type Queries struct {
	client *user.Client
}

func NewQueries(client *user.Client) port.UserQueries {
	return &Queries{client: client}
}

func (u *Queries) FindById(ctx context.Context, id string) (*port.UserDTO, error) {
	entDto, err := u.client.User.
		Query().
		Where(ent.ID(uuid.Parse(id)), ent.DeletedAtIsNil()).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return &port.UserDTO{
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
