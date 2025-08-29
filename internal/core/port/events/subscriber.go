package events

import (
	"context"
)

type Subscriber interface {
	Dispatch(context.Context, Event) error
}
