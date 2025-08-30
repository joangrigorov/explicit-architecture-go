package user

import (
	"fmt"
	"time"
)

type ID string

func (i ID) String() string {
	return string(i)
}

type IdPUserID string

func (i IdPUserID) String() string {
	return string(i)
}

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
	confirmedAt := time.Now()
	u.ConfirmedAt = &confirmedAt
	u.recordEvent(NewConfirmedEvent(u.ID, confirmedAt))
}

func (u *User) LinkToIdP(idpUserID IdPUserID) {
	u.IdPUserId = &idpUserID
	u.recordEvent(NewIdPUserLinkedEvent(u.ID, idpUserID))
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func NewUser(id ID, username Username, email Email, fName string, lName string, role Role) *User {
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
			NewCreatedEvent(id, username, email, fName, lName),
		},
	}
}
