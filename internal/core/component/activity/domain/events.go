package domain

import "time"

type ActivityEvent interface {
	Name() string
	CreatedAt() time.Time
}
type ActivityCreated struct {
	createdAt time.Time
}

func NewActivityCreated() *ActivityCreated {
	return &ActivityCreated{createdAt: time.Now()}
}

func (a *ActivityCreated) Name() string {
	return "Activity.ActivityCreated"
}

func (a *ActivityCreated) CreatedAt() time.Time {
	return a.createdAt
}
