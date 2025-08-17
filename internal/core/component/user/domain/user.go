package domain

import (
	"app/internal/core/shared_kernel/domain"
	"time"
)

type User struct {
	ID          domain.UserID
	Username    string
	Email       string
	FirstName   string
	LastName    string
	ConfirmedAt *time.Time
	Role        Role
	IdPUserId   *domain.IdPUserID

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Confirm() {
	now := time.Now()
	u.ConfirmedAt = &now
}

func NewUser(
	id domain.UserID,
	username string,
	email string,
	fName string,
	lName string,
	role Role,
) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		FirstName: fName,
		LastName:  lName,
		Role:      role,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ReconstituteUser(
	id domain.UserID,
	username string,
	email string,
	firstName string,
	lastName string,
	role Role,
	idpUserId *domain.IdPUserID,
	confirmedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		ID:          id,
		Username:    username,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		Role:        role,
		IdPUserId:   idpUserId,
		ConfirmedAt: confirmedAt,

		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
