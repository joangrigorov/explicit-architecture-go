package event_bus

import (
	eventBus "app/internal/core/shared_kernel/events"
	"app/mock/core/port/logging"
	. "app/mock/core/shared_kernel/events"
	. "app/mock/infrastructure/framework/event_bus"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewSimpleEventBus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert.NotNil(t, NewSimpleEventBus(logging.NewMockLogger(ctrl)))
}

func TestNewEventBus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert.NotNil(t, NewEventBus(NewSimpleEventBus(logging.NewMockLogger(ctrl))))
}

func TestSimpleEventBus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := logging.NewMockLogger(ctrl)

	event1 := NewMockEvent(ctrl)
	event1.EXPECT().ID().Return(eventBus.EventID("example-event")).AnyTimes()

	event2 := NewMockEvent(ctrl)
	event2.EXPECT().ID().Return(eventBus.EventID("example-event-2")).AnyTimes()

	sub1 := NewMockSubscriber(ctrl)
	sub2 := NewMockSubscriber(ctrl)

	bus := NewSimpleEventBus(logger)
	defer bus.Close()

	bus.Subscribe(sub1, event1)
	bus.Subscribe(sub2, event2)

	t.Run("success", func(t *testing.T) {
		sub1ch := make(chan eventBus.Event, 1)
		sub1.EXPECT().Dispatch(event1).Do(func(event eventBus.Event) {
			sub1ch <- event
		}).Return(nil)
		sub2.EXPECT().Dispatch(event2).Times(0) // publishing event1 won't trigger sub2

		logger.EXPECT().Debug("Published event example-event, handled by *events.MockSubscriber")
		logger.EXPECT().Debug("Dispatched event example-event, handled by *events.MockSubscriber")
		bus.Publish(event1)

		select {
		case e := <-sub1ch:
			assert.Equal(t, event1, e)
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for event")
		}
	})

	t.Run("recover from panic", func(t *testing.T) {
		sub1.EXPECT().
			Dispatch(event1).
			Do(func(event eventBus.Event) {
				panic("test panic")
			}).
			Times(1).
			Return(nil)
		sub2.EXPECT().Dispatch(event2).Times(0) // publishing event1 won't trigger sub2

		logger.EXPECT().Debug("Published event example-event, handled by *events.MockSubscriber").Times(2)
		logger.EXPECT().Error("Recovered from panic in subscriber *events.MockSubscriber: test panic").Times(1)
		bus.Publish(event1)

		sub1.EXPECT().Dispatch(event1).Return(errors.New("test error"))

		logger.EXPECT().Error(errors.New("test error")).Times(1)
		bus.Publish(event1)
	})
}
