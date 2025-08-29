package events

import (
	eventBus "app/internal/core/port/events"
	"app/mock/core/port/events"
	"app/mock/core/port/logging"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type fakeEvent1 struct{}

func (f fakeEvent1) CreatedAt() time.Time {
	return time.Time{}
}

type fakeEvent2 struct{}

func (f fakeEvent2) CreatedAt() time.Time {
	return time.Time{}
}

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

	ctx := context.Background()

	logger := logging.NewMockLogger(ctrl)

	event1 := fakeEvent1{}
	event2 := fakeEvent2{}

	sub1 := events.NewMockSubscriber(ctrl)
	sub2 := events.NewMockSubscriber(ctrl)

	bus := NewSimpleEventBus(logger)
	defer bus.Close()

	bus.Subscribe(sub1, event1)
	bus.Subscribe(sub2, event2)

	t.Run("success", func(t *testing.T) {
		sub1ch := make(chan eventBus.Event, 1)
		sub1.EXPECT().Dispatch(gomock.Any(), event1).Do(func(c context.Context, event eventBus.Event) {
			sub1ch <- event
		}).Return(nil)
		sub2.EXPECT().Dispatch(ctx, event2).Times(0) // publishing event1 won't trigger sub2

		logger.EXPECT().Debug("Published event app/internal/infrastructure/framework/events.fakeEvent1, handled by app/mock/core/port/events.MockSubscriber")
		logger.EXPECT().Debug("Handled event app/internal/infrastructure/framework/events.fakeEvent1 by app/mock/core/port/events.MockSubscriber")
		bus.Publish(ctx, event1)

		select {
		case e := <-sub1ch:
			assert.Equal(t, event1, e)
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for event")
		}
	})

	t.Run("recover from panic", func(t *testing.T) {
		sub1.EXPECT().
			Dispatch(gomock.Any(), event1).
			Do(func(c context.Context, event eventBus.Event) {
				panic("test panic")
			}).
			Times(1).
			Return(nil)
		sub2.EXPECT().Dispatch(gomock.Any(), event2).Times(0) // publishing event1 won't trigger sub2

		logger.EXPECT().Debug("Published event app/internal/infrastructure/framework/events.fakeEvent1, handled by app/mock/core/port/events.MockSubscriber").Times(1)
		logger.EXPECT().Error("Recovered from panic in subscriber app/mock/core/port/events.MockSubscriber: test panic").Times(1)
		bus.Publish(ctx, event1)
	})

	t.Run("recover from error", func(t *testing.T) {
		sub1.EXPECT().Dispatch(gomock.Any(), event1).Return(errors.New("test error"))
		sub2.EXPECT().Dispatch(gomock.Any(), event2).Times(0) // publishing event1 won't trigger sub2

		logger.EXPECT().Debug("Published event app/internal/infrastructure/framework/events.fakeEvent1, handled by app/mock/core/port/events.MockSubscriber").Times(1)
		logger.EXPECT().Error(errors.New("test error")).Times(1)
		bus.Publish(ctx, event1)
	})
}
