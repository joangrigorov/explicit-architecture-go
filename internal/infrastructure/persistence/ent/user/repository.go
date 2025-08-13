package user

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/persistence/ent/generated/user"
	ent "app/internal/infrastructure/persistence/ent/generated/user/user"
	"context"
)

type Repository struct {
	client *user.Client
}

func (r *Repository) GetByIdPUserId(ctx context.Context, idPUserId domain.IdPUserId) (*domain.User, error) {
	dto, err := r.client.User.
		Query().
		Where(
			ent.IdpUserID(string(idPUserId)),
			ent.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	entity := mapEntity(dto)

	return entity, nil
}

func (r *Repository) Create(ctx context.Context, u *domain.User) error {
	_, err := r.client.User.
		Create().
		SetID(uuid.Parse(u.Id)).
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetRole(mapDtoRole(u.Role)).
		SetNillableIdpUserID((*string)(u.IdPUserId)).
		SetNillableConfirmedAt(u.ConfirmedAt).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(u.UpdatedAt).
		Save(ctx)

	return err
}

func NewRepository(c *user.Client) repositories.UserRepository {
	return &Repository{client: c}
}
