package events

import (
	. "app/internal/core/port/events"
	"app/internal/core/port/logging"
	. "app/internal/core/shared_kernel/events"
	"fmt"
	"reflect"
	"sync"
)

// bufferSize of how many events can wait in the queue before they are processed
const bufferSize = 100

type subscriberWorker struct {
	sub Subscriber
	ch  chan Event
}

type SimpleEventBus struct {
	logger logging.Logger
	subs   map[EventID][]*subscriberWorker
	mu     sync.RWMutex
	wg     sync.WaitGroup
}

func NewSimpleEventBus(logger logging.Logger) *SimpleEventBus {
	return &SimpleEventBus{
		logger: logger,
		subs:   make(map[EventID][]*subscriberWorker),
	}
}

func NewEventBus(bus *SimpleEventBus) EventBus {
	return bus
}

// Subscribe to event
func (b *SimpleEventBus) Subscribe(subscriber Subscriber, event Event) {
	b.mu.Lock()
	defer b.mu.Unlock()

	worker := &subscriberWorker{
		sub: subscriber,
		ch:  make(chan Event, bufferSize), // configurable buffer size. TODO use configuration value for buffer size
	}

	b.wg.Add(1)
	go func(w *subscriberWorker) {
		defer b.wg.Done()

		for e := range w.ch {
			// Wrap each dispatch in its own func scope
			func(ev Event) {
				defer func() {
					if r := recover(); r != nil {
						b.logger.Error(fmt.Sprintf(
							"Recovered from panic in subscriber %s: %v",
							reflect.TypeOf(w.sub).String(), r,
						))
					}
				}()

				if err := w.sub.Dispatch(ev); err != nil {
					b.logger.Error(err)
				} else {
					b.logger.Debug(fmt.Sprintf(
						"Dispatched event %s, handled by %s",
						ev.ID(), reflect.TypeOf(w.sub).String(),
					))
				}
			}(e)
		}
	}(worker)

	b.subs[event.ID()] = append(b.subs[event.ID()], worker)
}

// Publish sends to subscriber channels (non-blocking, log drops)
func (b *SimpleEventBus) Publish(event Event) error {
	b.mu.RLock()
	subs := append([]*subscriberWorker(nil), b.subs[event.ID()]...) // copy slice
	b.mu.RUnlock()

	for _, w := range subs {
		select {
		case w.ch <- event:
			b.logger.Debug(fmt.Sprintf(
				"Published event %s, handled by %s",
				event.ID(),
				reflect.TypeOf(w.sub).String(),
			))
		default: // coverage-ignore
			b.logger.Error(fmt.Sprintf(
				"Dropped event %s for subscriber %s",
				event.ID(),
				reflect.TypeOf(w.sub).String(),
			))
		}
	}

	return nil
}

// Close gracefully shuts down all workers
func (b *SimpleEventBus) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, workers := range b.subs {
		for _, w := range workers {
			close(w.ch)
		}
	}
	b.wg.Wait()
}
