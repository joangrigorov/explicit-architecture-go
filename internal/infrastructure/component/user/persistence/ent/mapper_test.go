package ent

import (
	"app/internal/core/component/user/domain"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	roles "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"app/internal/infrastructure/framework/uuid"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestMapDomainRole(t *testing.T) {
	t.Run("supported", func(t *testing.T) {
		assert.Equal(t, &domain.Admin{}, mapDomainRole(roles.RoleAdmin))
		assert.Equal(t, &domain.Member{}, mapDomainRole(roles.RoleMember))
	})

	t.Run("panics on unsupported", func(t *testing.T) {
		assert.Panics(t, func() {
			const UnsupportedRole roles.Role = "unsupported"
			mapDomainRole(UnsupportedRole)
		})
	})
}

type unsupportedRole struct{}

func (u *unsupportedRole) ID() domain.RoleId {
	return "unsupported"
}

func (u *unsupportedRole) String() string {
	return "unsupported"
}

func TestMapDtoRole(t *testing.T) {
	t.Run("supported", func(t *testing.T) {
		assert.Equal(t, roles.RoleAdmin, mapDtoRole(&domain.Admin{}))
		assert.Equal(t, roles.RoleMember, mapDtoRole(&domain.Member{}))
	})

	t.Run("panics on unsupported", func(t *testing.T) {
		assert.Panics(t, func() {
			mapDtoRole(&unsupportedRole{})
		})
	})
}

func TestMapEntity(t *testing.T) {
	f := faker.New()

	id := f.UUID().V4()
	username := f.Internet().User()
	email := f.Internet().Email()
	fName := f.Person().FirstName()
	lName := f.Person().LastName()
	role := f.RandomStringElement([]string{"admin", "member"})
	idPUserId := f.UUID().V4()
	confirmedAt := f.Time().Time(time.Now())
	createdAt := f.Time().Time(time.Now())
	updatedAt := f.Time().Time(time.Now())

	entity := mapUserAggregate(&generated.User{
		ID:          uuid.Parse(id),
		Username:    username,
		Email:       email,
		FirstName:   fName,
		LastName:    lName,
		Role:        roles.Role(role),
		IdpUserID:   &idPUserId,
		ConfirmedAt: &confirmedAt,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	})

	assert.Equal(t, id, string(entity.ID))
	assert.Equal(t, username, entity.Username)
	assert.Equal(t, email, entity.Email)
	assert.Equal(t, fName, entity.FirstName)
	assert.Equal(t, lName, entity.LastName)
	assert.Equal(t, role, entity.Role.ID().String())
	assert.Equal(t, idPUserId, string(*entity.IdPUserId))
	assert.Equal(t, &confirmedAt, entity.ConfirmedAt)
	assert.Equal(t, createdAt, entity.CreatedAt)
	assert.Equal(t, updatedAt, entity.UpdatedAt)
}
