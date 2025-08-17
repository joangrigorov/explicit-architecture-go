package events

import (
	eventBus "app/internal/core/port/events"
	"app/internal/core/shared_kernel/events"
)

// TransactionalEventBus collects events and publishes them in a batch (see TransactionalEventBus.Flush())
type TransactionalEventBus struct {
	outbox   []events.Event
	eventBus eventBus.EventBus
}

func NewTransactionalEventBus(bus *SimpleEventBus) *TransactionalEventBus {
	return &TransactionalEventBus{eventBus: bus}
}

func (t *TransactionalEventBus) Publish(event events.Event) error {
	t.outbox = append(t.outbox, event)
	return nil
}

func (t *TransactionalEventBus) Flush() error {
	for _, event := range t.outbox {
		if err := t.eventBus.Publish(event); err != nil {
			return err
		}
	}
	t.Reset()
	return nil
}

func (t *TransactionalEventBus) Reset() {
	t.outbox = []events.Event{}
}
