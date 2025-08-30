package user

import (
	"app/internal/core/shared_kernel/domain/mixin"
	"time"
)

type Event interface {
	UserID() ID
	CreatedAt() time.Time
}

type CreatedEvent struct {
	userId    ID
	username  Username
	email     Email
	firstName string
	lastName  string

	mixin.WithCreatedAt
}

type IdPUserLinkedEvent struct {
	userID    ID
	idpUserID IdPUserID

	mixin.WithCreatedAt
}

type ConfirmedEvent struct {
	userID      ID
	confirmedAt time.Time

	mixin.WithCreatedAt
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

func NewCreatedEvent(id ID, username Username, email Email, fName string, lName string) CreatedEvent {
	return CreatedEvent{
		userId:        id,
		username:      username,
		email:         email,
		firstName:     fName,
		lastName:      lName,
		WithCreatedAt: mixin.NewWithCreatedAtNow(),
	}
}

func NewIdPUserLinkedEvent(userID ID, idpUserID IdPUserID) IdPUserLinkedEvent {
	return IdPUserLinkedEvent{
		userID:        userID,
		idpUserID:     idpUserID,
		WithCreatedAt: mixin.NewWithCreatedAtNow(),
	}
}

func NewConfirmedEvent(userId ID, confirmedAt time.Time) ConfirmedEvent {
	return ConfirmedEvent{
		userID:        userId,
		confirmedAt:   confirmedAt,
		WithCreatedAt: mixin.NewWithCreatedAtNow(),
	}
}
