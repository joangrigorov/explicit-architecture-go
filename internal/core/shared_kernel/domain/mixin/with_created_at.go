package mixin

import "time"

type WithCreatedAt struct {
	createdAt time.Time
}

func NewWithCreatedAt(createdAt time.Time) WithCreatedAt {
	return WithCreatedAt{createdAt: createdAt}
}

func NewWithCreatedAtNow() WithCreatedAt {
	return NewWithCreatedAt(time.Now())
}

func (w WithCreatedAt) CreatedAt() time.Time {
	return w.createdAt
}
