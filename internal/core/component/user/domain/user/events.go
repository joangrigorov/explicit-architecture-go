package user

import (
	"app/internal/core/shared_kernel/domain"
	"time"
)

type Event interface {
	UserID() ID
	CreatedAt() time.Time
}

type CreatedEvent struct {
	userId   ID
	username Username
	email    Email

	domain.WithCreatedAt
}

type IdPUserLinkedEvent struct {
	userID  ID
	idpUser IdPUserID

	domain.WithCreatedAt
}

type ConfirmedEvent struct {
	userID ID

	domain.WithCreatedAt
}

func (u CreatedEvent) UserID() ID {
	return u.userId
}
func (u CreatedEvent) Email() Email {
	return u.email
}
func (u CreatedEvent) Username() Username {
	return u.username
}

func (i IdPUserLinkedEvent) UserID() ID {
	return i.userID
}

func (u ConfirmedEvent) UserID() ID {
	return u.userID
}

func NewIdPUserLinkedEvent(userID ID, idpUserID IdPUserID) IdPUserLinkedEvent {
	return IdPUserLinkedEvent{
		userID:        userID,
		idpUser:       idpUserID,
		WithCreatedAt: domain.NewWithCreatedAtNow(),
	}
}

func NewConfirmedEvent(userId ID) ConfirmedEvent {
	return ConfirmedEvent{
		userID:        userId,
		WithCreatedAt: domain.NewWithCreatedAtNow(),
	}
}

func NewCreatedEvent(id ID, username Username, email Email) CreatedEvent {
	return CreatedEvent{
		userId:        id,
		username:      username,
		email:         email,
		WithCreatedAt: domain.NewWithCreatedAtNow(),
	}
}
