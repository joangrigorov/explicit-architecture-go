package events

import (
	"app/internal/core/shared_kernel/domain"
)

type UserCreated struct {
	userId   domain.UserID
	email    string
	password string
}

func (u UserCreated) ID() EventID {
	return makeEventID(u)
}

func NewUserCreated(id domain.UserID, email string, password string) UserCreated {
	return UserCreated{
		userId:   id,
		email:    email,
		password: password,
	}
}
