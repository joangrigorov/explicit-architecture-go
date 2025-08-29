package events

import (
	eventBus "app/internal/core/port/events"
	"context"
)

type outboxItem struct {
	event eventBus.Event
	ctx   context.Context
}

// TransactionalEventBus collects events and publishes them in a batch (see TransactionalEventBus.Flush())
type TransactionalEventBus struct {
	outbox   []*outboxItem
	eventBus eventBus.EventBus
}

func NewTransactionalEventBus(bus *SimpleEventBus) *TransactionalEventBus {
	return &TransactionalEventBus{eventBus: bus}
}

func (t *TransactionalEventBus) Publish(ctx context.Context, event eventBus.Event) error {
	t.outbox = append(t.outbox, &outboxItem{ctx: ctx, event: event})
	return nil
}

func (t *TransactionalEventBus) Flush() error {
	for _, item := range t.outbox {
		if err := t.eventBus.Publish(item.ctx, item.event); err != nil {
			return err
		}
	}
	t.Reset()
	return nil
}

func (t *TransactionalEventBus) Reset() {
	t.outbox = make([]*outboxItem, 0)
}

func CloseEventBus(ctx context.Context, bus *TransactionalEventBus, err *error) {
	if r := recover(); r != nil {
		bus.Reset()
		panic(r)
	} else if *err != nil || ctx.Err() != nil {
		// If there was an error or context was canceled
		bus.Reset()
	} else {
		*err = bus.Flush()
	}
}
