package domain

import (
	"app/internal/core/shared_kernel/domain"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	f := faker.New()
	userId := domain.UserID(f.UUID().V4())
	username := f.Internet().User()
	email := f.Internet().Email()
	fName := f.Person().FirstName()
	lName := f.Person().LastName()
	role := Admin{}

	user := NewUser(userId, username, email, fName, lName, role)

	assert.NotNil(t, user)
	assert.Equal(t, userId, user.ID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, fName, user.FirstName)
	assert.Equal(t, lName, user.LastName)
	assert.Equal(t, role, user.Role)
	assert.Nil(t, user.ConfirmedAt)
	assert.Nil(t, user.IdPUserId)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
}

func TestReconstituteUser(t *testing.T) {
	f := faker.New()
	userId := domain.UserID(f.UUID().V4())
	username := f.Internet().User()
	email := f.Internet().Email()
	fName := f.Person().FirstName()
	lName := f.Person().LastName()
	role := Admin{}
	idPUserId := domain.IdPUserID(f.UUID().V4())
	createdAt := f.Time().Time(time.Now())
	updatedAt := f.Time().Time(time.Now())
	confirmedAt := f.Time().Time(time.Now())

	user := ReconstituteUser(
		userId,
		username,
		email,
		fName,
		lName,
		role,
		&idPUserId,
		&confirmedAt,
		createdAt,
		updatedAt,
	)

	assert.NotNil(t, user)
	assert.Equal(t, userId, user.ID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, fName, user.FirstName)
	assert.Equal(t, lName, user.LastName)
	assert.Equal(t, role, user.Role)
	assert.Equal(t, &confirmedAt, user.ConfirmedAt)
	assert.Equal(t, createdAt, user.CreatedAt)
	assert.Equal(t, updatedAt, user.UpdatedAt)
}

func TestUser_Confirm(t *testing.T) {
	user := &User{}

	assert.Nil(t, user.ConfirmedAt)
	user.Confirm()
	assert.NotNil(t, user.ConfirmedAt)
}
