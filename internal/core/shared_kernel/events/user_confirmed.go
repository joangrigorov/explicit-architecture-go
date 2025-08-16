package events

import (
	"app/internal/core/shared_kernel/domain"
)

type UserConfirmed struct {
	userId domain.UserID
}

func (u UserConfirmed) ID() EventID {
	return makeEventID(u)
}

func NewUserConfirmed(userId domain.UserID) UserConfirmed {
	return UserConfirmed{userId: userId}
}
