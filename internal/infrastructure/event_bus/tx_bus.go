package event_bus

import (
	eventBus "app/internal/core/port/event_bus"
	"app/internal/core/shared_kernel/events"
	"context"
)

type outboxItem struct {
	event events.Event
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

func (t *TransactionalEventBus) Publish(ctx context.Context, event events.Event) error {
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
