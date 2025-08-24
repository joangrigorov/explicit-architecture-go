package ent

import (
	"app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"app/internal/infrastructure/framework/uuid"
	"context"
	"time"
)

type UserRepository struct {
	entTx     *generated.Tx
	entClient *generated.Client
}

func (r *UserRepository) Update(ctx context.Context, u *User) error {
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

func (r *UserRepository) GetById(ctx context.Context, id UserID) (*User, error) {
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

	entity := mapUserAggregate(dto)

	return entity, nil
}

func (r *UserRepository) GetByIdPUserId(ctx context.Context, idPUserId IdPUserID) (*User, error) {
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

	entity := mapUserAggregate(dto)

	return entity, nil
}

func (r *UserRepository) Create(ctx context.Context, u *User) error {
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

func (r *UserRepository) client() *generated.UserClient {
	if r.entTx != nil {
		return r.entTx.User
	}

	if r.entClient != nil {
		return r.entClient.User
	}

	panic("ent client not initialized")
}

func (r *UserRepository) WithTx(tx *generated.Tx) *UserRepository {
	return &UserRepository{entTx: tx, entClient: r.entClient}
}

func NewRepository(r *UserRepository) repositories.UserRepository {
	return r
}
func NewConcreteRepository(c *generated.Client) *UserRepository {
	return &UserRepository{entClient: c}
}
