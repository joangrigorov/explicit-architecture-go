package user

import (
	"fmt"
	"time"
)

type User struct {
	ID          ID
	Username    Username
	Email       Email
	FirstName   string
	LastName    string
	ConfirmedAt *time.Time
	Role        Role
	IdPUserId   *IdPUserID

	CreatedAt time.Time
	UpdatedAt time.Time

	events []Event
}

func (u *User) recordEvent(event Event) {
	u.events = append(u.events, event)
}

func (u *User) ResetEvents() {
	u.events = make([]Event, 0)
}

func (u *User) Events() []Event {
	return u.events
}

func (u *User) Confirm() {
	now := time.Now()
	u.ConfirmedAt = &now
	u.recordEvent(NewConfirmedEvent(u.ID))
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) LinkToIdP(idpUserID IdPUserID) {
	u.IdPUserId = &idpUserID
	u.recordEvent(NewIdPUserLinkedEvent(u.ID, idpUserID))
}

func NewUser(
	id ID,
	username Username,
	email Email,
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

		events: []Event{
			NewCreatedEvent(id, username, email),
		},
	}
}

func ReconstituteUser(
	id ID,
	username Username,
	email Email,
	firstName string,
	lastName string,
	role Role,
	idpUserId *IdPUserID,
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
