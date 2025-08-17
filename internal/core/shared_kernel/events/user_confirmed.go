package events

import (
	"app/internal/core/shared_kernel/domain"
)

type UserConfirmed struct {
	userID domain.UserID
}

func (u UserConfirmed) ID() EventID {
	return makeEventID(u)
}

func (u UserConfirmed) UserID() domain.UserID {
	return u.userID
}

func NewUserConfirmed(userId domain.UserID) UserConfirmed {
	return UserConfirmed{userID: userId}
}
