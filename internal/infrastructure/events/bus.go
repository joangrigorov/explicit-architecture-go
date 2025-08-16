package events

import (
	. "app/internal/core/port/events"
	. "app/internal/core/shared_kernel/events"
	"sync"
)

type SimpleEventBus struct {
	subs map[EventID][]Subscriber
	wg   sync.WaitGroup
	mu   sync.RWMutex
}

func NewSimpleEventBus() *SimpleEventBus {
	return &SimpleEventBus{
		subs: make(map[EventID][]Subscriber),
	}
}

func NewEventBus() EventBus {
	return NewSimpleEventBus()
}

func (b *SimpleEventBus) Subscribe(subscriber Subscriber, event Event) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.subs[event.ID()] = append(b.subs[event.ID()], subscriber)
}

func (b *SimpleEventBus) Publish(event Event) {
	b.mu.RLock()
	subs := append([]Subscriber(nil), b.subs[event.ID()]...) // copy slice
	b.mu.RUnlock()

	for _, s := range subs {
		b.wg.Add(1)
		go b.dispatch(s, event)
	}
}

func (b *SimpleEventBus) dispatch(s Subscriber, e Event) {
	defer b.wg.Done()

	if err := s.Dispatch(e); err != nil { // coverage-ignore
		// TODO: log error
	}
}

func (b *SimpleEventBus) Wait() {
	b.wg.Wait()
}
