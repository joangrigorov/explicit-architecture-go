package events

import (
	. "app/internal/core/port/events"
	"app/internal/core/port/logging"
	"context"
	"fmt"
	"reflect"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// bufferSize of how many events can wait in the queue before they are processed
const bufferSize = 100

type event struct {
	e   Event
	ctx context.Context
}

type subscriberWorker struct {
	sub Subscriber
	ch  chan event
}

type SimpleEventBus struct {
	logger logging.Logger
	tracer trace.Tracer

	subs map[string][]*subscriberWorker
	mu   sync.RWMutex
	wg   sync.WaitGroup
}

func NewSimpleEventBus(logger logging.Logger, tracer trace.Tracer) *SimpleEventBus {
	return &SimpleEventBus{
		logger: logger,
		subs:   make(map[string][]*subscriberWorker),
		tracer: tracer,
	}
}

func NewEventBus(bus *SimpleEventBus) EventBus {
	return bus
}

// Subscribe to event
func (b *SimpleEventBus) Subscribe(subscriber Subscriber, e Event) {
	b.mu.Lock()
	defer b.mu.Unlock()

	worker := &subscriberWorker{
		sub: subscriber,
		ch:  make(chan event, bufferSize), // configurable buffer size. TODO use configuration value for buffer size
	}

	b.wg.Add(1)
	go func(w *subscriberWorker) {
		defer b.wg.Done()

		for ev := range w.ch {
			// Wrap each dispatch in its own func scope
			func(ev event) {
				ctx, span := b.tracer.Start(ev.ctx, fmt.Sprintf("Handle event %T", ev.e))
				defer span.End()
				span.SetAttributes(
					attribute.String("event", id(ev.e)),
					attribute.String("sub", id(w.sub)),
				)

				defer func() {
					if r := recover(); r != nil {
						b.logger.Error(fmt.Sprintf("Recovered from panic in subscriber %s: %v", id(w.sub), r))
					}
				}()

				if err := w.sub.Dispatch(ctx, ev.e); err != nil {
					b.logger.Error(err)
					span.RecordError(err)
				} else {
					span.AddEvent(fmt.Sprintf("Handled event %s by %s", id(ev.e), id(w.sub)))
					b.logger.Debug(fmt.Sprintf("Handled event %s by %s", id(ev.e), id(w.sub)))
				}
			}(ev)
		}
	}(worker)

	eventID := id(e)
	b.subs[eventID] = append(b.subs[eventID], worker)
}

// Publish sends to subscriber channels (non-blocking, log drops)
func (b *SimpleEventBus) Publish(ctx context.Context, e Event) error {
	b.mu.RLock()
	eventID := id(e)
	subs := append([]*subscriberWorker(nil), b.subs[eventID]...) // copy slice
	b.mu.RUnlock()

	for _, w := range subs {
		subID := id(w.sub)
		select {
		case w.ch <- event{ctx: context.WithoutCancel(ctx), e: e}:
			b.logger.Debug(fmt.Sprintf("Published event %s, handled by %s", eventID, subID))
		default: // coverage-ignore
			b.logger.Error(fmt.Sprintf("Dropped event %s for subscriber %s", eventID, subID))
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

func id(e interface{}) string {
	t := reflect.TypeOf(e)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.PkgPath() + "." + t.Name()
}
