package events

import "time"

type Event interface {
	CreatedAt() time.Time
}
