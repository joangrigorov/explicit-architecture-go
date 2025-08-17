package events

import (
	"app/internal/core/shared_kernel/domain"
)

type UserCreated struct {
	userId   domain.UserID
	username string
	email    string
	password string
}

func (u UserCreated) UserId() domain.UserID {
	return u.userId
}

func (u UserCreated) Email() string {
	return u.email
}

func (u UserCreated) Password() string {
	return u.password
}

func (u UserCreated) ID() EventID {
	return makeEventID(u)
}

func (u UserCreated) Username() string {
	return u.username
}

func NewUserCreated(id domain.UserID, username string, email string, password string) UserCreated {
	return UserCreated{
		userId:   id,
		username: username,
		email:    email,
		password: password,
	}
}
