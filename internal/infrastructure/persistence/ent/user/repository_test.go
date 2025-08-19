package user

import (
	"app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/infrastructure/persistence/ent/generated/user/enttest"
	"app/internal/infrastructure/persistence/ent/generated/user/user"
	"app/internal/infrastructure/uuid"
	"context"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	assert.NotNil(t, NewRepository(nil))
}

func TestRepository_Create(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	id,
		username,
		email,
		fName,
		lName,
		role,
		idPUserId,
		confirmedAt,
		createdAt,
		updatedAt := fakeUserData(f)

	u := &domain.User{
		ID:          UserID(id),
		Username:    username,
		Email:       email,
		FirstName:   fName,
		LastName:    lName,
		ConfirmedAt: &confirmedAt,
		Role:        mapDomainRole(user.Role(role)),
		IdPUserId:   &idPUserId,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	r := NewRepository(client)

	assert.NoError(t, r.Create(ctx, u))
}

func TestRepository_Update(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	id,
		username,
		email,
		fName,
		lName,
		role,
		idPUserId,
		confirmedAt,
		createdAt,
		updatedAt := fakeUserData(f)

	_, err := client.User.Create().
		SetID(uuid.Parse(id)).
		SetUsername(username).
		SetEmail(email).
		SetFirstName(fName).
		SetLastName(lName).
		SetRole(user.Role(role)).
		SetIdpUserID(string(idPUserId)).
		SetConfirmedAt(confirmedAt).
		SetCreatedAt(createdAt).
		SetUpdatedAt(updatedAt).
		Save(ctx)

	assert.NoError(t, err)

	u := &domain.User{
		ID:          UserID(id),
		Username:    username,
		Email:       email,
		FirstName:   fName,
		LastName:    lName,
		ConfirmedAt: &confirmedAt,
		Role:        mapDomainRole(user.Role(role)),
		IdPUserId:   &idPUserId,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	r := NewRepository(client)

	assert.NoError(t, r.Update(ctx, u))
	assert.NotEqual(t, updatedAt, u.CreatedAt)
}

func TestRepository_GetByIdPUserId(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	t.Run("exists", func(t *testing.T) {
		id,
			username,
			email,
			fName,
			lName,
			role,
			idPUserId,
			confirmedAt,
			createdAt,
			updatedAt := fakeUserData(f)

		_, err := client.User.Create().
			SetID(uuid.Parse(id)).
			SetUsername(username).
			SetEmail(email).
			SetFirstName(fName).
			SetLastName(lName).
			SetRole(user.Role(role)).
			SetIdpUserID(string(idPUserId)).
			SetConfirmedAt(confirmedAt).
			SetCreatedAt(createdAt).
			SetUpdatedAt(updatedAt).
			Save(ctx)

		assert.NoError(t, err)

		r := &Repository{entClient: client}
		u, err := r.GetByIdPUserId(ctx, idPUserId)

		assert.NoError(t, err)
		assert.NotNil(t, u)

		assert.Equal(t, id, string(u.ID))
		assert.Equal(t, email, u.Email)
		assert.Equal(t, username, u.Username)
		assert.Equal(t, email, u.Email)
		assert.Equal(t, fName, u.FirstName)
		assert.Equal(t, lName, u.LastName)
		assert.Equal(t, role, u.Role.ID().String())
		assert.Equal(t, &idPUserId, u.IdPUserId)
		assert.Equal(t, confirmedAt, u.ConfirmedAt.In(time.Local))
		assert.Equal(t, createdAt, u.CreatedAt.In(time.Local))
		assert.Equal(t, updatedAt, u.UpdatedAt.In(time.Local))
	})

	t.Run("soft-deleted", func(t *testing.T) {
		id,
			username,
			email,
			fName,
			lName,
			role,
			idPUserId,
			confirmedAt,
			createdAt,
			updatedAt := fakeUserData(f)

		_, err := client.User.Create().
			SetID(uuid.Parse(id)).
			SetUsername(username).
			SetEmail(email).
			SetFirstName(fName).
			SetLastName(lName).
			SetRole(user.Role(role)).
			SetIdpUserID(string(idPUserId)).
			SetConfirmedAt(confirmedAt).
			SetCreatedAt(createdAt).
			SetUpdatedAt(updatedAt).
			SetDeletedAt(time.Now()).
			Save(ctx)

		assert.NoError(t, err)

		r := &Repository{entClient: client}
		u, err := r.GetByIdPUserId(ctx, idPUserId)

		assert.Error(t, err, "user: user not found")
		assert.Nil(t, u)
	})
}

func TestRepository_GetById(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	t.Run("exists", func(t *testing.T) {
		id,
			username,
			email,
			fName,
			lName,
			role,
			idPUserId,
			confirmedAt,
			createdAt,
			updatedAt := fakeUserData(f)

		_, err := client.User.Create().
			SetID(uuid.Parse(id)).
			SetUsername(username).
			SetEmail(email).
			SetFirstName(fName).
			SetLastName(lName).
			SetRole(user.Role(role)).
			SetIdpUserID(string(idPUserId)).
			SetConfirmedAt(confirmedAt).
			SetCreatedAt(createdAt).
			SetUpdatedAt(updatedAt).
			Save(ctx)

		assert.NoError(t, err)

		r := &Repository{entClient: client}
		u, err := r.GetById(ctx, UserID(id))

		assert.NoError(t, err)
		assert.NotNil(t, u)

		assert.Equal(t, id, string(u.ID))
		assert.Equal(t, email, u.Email)
		assert.Equal(t, username, u.Username)
		assert.Equal(t, email, u.Email)
		assert.Equal(t, fName, u.FirstName)
		assert.Equal(t, lName, u.LastName)
		assert.Equal(t, role, u.Role.ID().String())
		assert.Equal(t, &idPUserId, u.IdPUserId)
		assert.Equal(t, confirmedAt, u.ConfirmedAt.In(time.Local))
		assert.Equal(t, createdAt, u.CreatedAt.In(time.Local))
		assert.Equal(t, updatedAt, u.UpdatedAt.In(time.Local))
	})

	t.Run("soft-deleted", func(t *testing.T) {
		id,
			username,
			email,
			fName,
			lName,
			role,
			idPUserId,
			confirmedAt,
			createdAt,
			updatedAt := fakeUserData(f)

		_, err := client.User.Create().
			SetID(uuid.Parse(id)).
			SetUsername(username).
			SetEmail(email).
			SetFirstName(fName).
			SetLastName(lName).
			SetRole(user.Role(role)).
			SetIdpUserID(string(idPUserId)).
			SetConfirmedAt(confirmedAt).
			SetCreatedAt(createdAt).
			SetUpdatedAt(updatedAt).
			SetDeletedAt(time.Now()).
			Save(ctx)

		assert.NoError(t, err)

		r := &Repository{entClient: client}
		u, err := r.GetById(ctx, UserID(id))

		assert.Error(t, err, "user: user not found")
		assert.Nil(t, u)
	})
}

func fakeUserData(f faker.Faker) (string, string, string, string, string, string, IdPUserID, time.Time, time.Time, time.Time) {
	id := f.UUID().V4()
	username := f.Internet().User()
	email := f.Internet().Email()
	fName := f.Person().FirstName()
	lName := f.Person().LastName()
	role := f.RandomStringElement([]string{"admin", "member"})
	idPUserId := IdPUserID(f.UUID().V4())
	confirmedAt := f.Time().Time(time.Now())
	createdAt := f.Time().Time(time.Now())
	updatedAt := f.Time().Time(time.Now())
	return id, username, email, fName, lName, role, idPUserId, confirmedAt, createdAt, updatedAt
}
