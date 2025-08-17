package user

import (
	. "app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/persistence/ent/generated/user"
	ent "app/internal/infrastructure/persistence/ent/generated/user/user"
	"context"
	"time"
)

type Repository struct {
	entTx     *user.Tx
	entClient *user.Client
}

func (r *Repository) Update(ctx context.Context, u *User) error {
	updatedAt := time.Now()
	_, err := r.
		client().
		UpdateOneID(uuid.Parse(u.ID)).
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetRole(mapDtoRole(u.Role)).
		SetNillableIdpUserID((*string)(u.IdPUserId)).
		SetNillableConfirmedAt(u.ConfirmedAt).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(updatedAt).
		Save(ctx)

	u.UpdatedAt = updatedAt

	return err
}

func (r *Repository) GetById(ctx context.Context, id UserID) (*User, error) {
	dto, err := r.
		client().
		Query().
		Where(
			ent.ID(uuid.Parse(id)),
			ent.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	entity := mapEntity(dto)

	return entity, nil
}

func (r *Repository) GetByIdPUserId(ctx context.Context, idPUserId IdPUserID) (*User, error) {
	dto, err := r.
		client().
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

func (r *Repository) Create(ctx context.Context, u *User) error {
	_, err := r.
		client().
		Create().
		SetID(uuid.Parse(u.ID)).
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

func (r *Repository) client() *user.UserClient {
	if r.entTx != nil {
		return r.entTx.User
	}

	if r.entClient != nil {
		return r.entClient.User
	}

	panic("ent client not initialized")
}

func (r *Repository) WithTx(tx *user.Tx) *Repository {
	return &Repository{entTx: tx, entClient: r.entClient}
}

func NewRepository(r *Repository) UserRepository {
	return r
}
func NewConcreteRepository(c *user.Client) *Repository {
	return &Repository{entClient: c}
}
