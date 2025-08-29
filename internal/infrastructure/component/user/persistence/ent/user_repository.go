package ent

import (
	"app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain/user"
	"app/internal/core/port/events"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	ent "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"app/internal/infrastructure/framework/uuid"
	"context"
	"time"
)

type UserRepository struct {
	eventBus  events.EventBus
	entTx     *generated.Tx
	entClient *generated.Client
}

func (r *UserRepository) Update(ctx context.Context, u *User) error {
	updatedAt := time.Now()
	_, err := r.
		client().
		UpdateOneID(uuid.Parse(u.ID)).
		SetUsername(u.Username.String()).
		SetEmail(u.Email.String()).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetRole(mapDtoRole(u.Role)).
		SetNillableIdpUserID((*string)(u.IdPUserId)).
		SetNillableConfirmedAt(u.ConfirmedAt).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(updatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	u.UpdatedAt = updatedAt

	if err = r.flushEvents(ctx, u); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, id ID) (*User, error) {
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

func (r *UserRepository) GetByEmail(ctx context.Context, email Email) (*User, error) {
	dto, err := r.
		client().
		Query().
		Where(
			ent.Email(email.String()),
			ent.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	entity := mapUserAggregate(dto)

	return entity, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username Username) (*User, error) {
	dto, err := r.
		client().
		Query().
		Where(
			ent.Username(username.String()),
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
		SetUsername(u.Username.String()).
		SetEmail(u.Email.String()).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetRole(mapDtoRole(u.Role)).
		SetNillableIdpUserID((*string)(u.IdPUserId)).
		SetNillableConfirmedAt(u.ConfirmedAt).
		SetCreatedAt(u.CreatedAt).
		SetUpdatedAt(u.UpdatedAt).
		Save(ctx)

	if err != nil {
		return err
	}

	if err = r.flushEvents(ctx, u); err != nil {
		return err
	}

	return nil
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

func (r *UserRepository) flushEvents(ctx context.Context, u *User) error {
	for _, event := range u.Events() {
		err := r.eventBus.Publish(ctx, event)
		if err != nil {
			return err
		}
	}

	u.ResetEvents()
	return nil
}

func (r *UserRepository) WithTx(tx *generated.Tx) *UserRepository {
	return &UserRepository{entTx: tx, entClient: r.entClient, eventBus: r.eventBus}
}

func (r *UserRepository) WithEventBus(eventBus events.EventBus) *UserRepository {
	return &UserRepository{entTx: r.entTx, entClient: r.entClient, eventBus: eventBus}
}

func NewRepository(r *UserRepository) repositories.UserRepository {
	return r
}
func NewConcreteRepository(c *generated.Client, eventBus events.EventBus) *UserRepository {
	return &UserRepository{entClient: c, eventBus: eventBus}
}
