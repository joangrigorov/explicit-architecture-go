package events

import (
	eventBus "app/internal/core/shared_kernel/events"
	. "app/mock/core/shared_kernel/events"
	. "app/mock/infrastructure/events"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewSimpleEventBus(t *testing.T) {
	assert.NotNil(t, NewSimpleEventBus())
}

func TestNewEventBus(t *testing.T) {
	assert.NotNil(t, NewEventBus())
}

func TestSimpleEventBus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	event1 := NewMockEvent(ctrl)
	event1.EXPECT().ID().Return(eventBus.EventID("example-event")).AnyTimes()

	event2 := NewMockEvent(ctrl)
	event2.EXPECT().ID().Return(eventBus.EventID("example-event-2")).AnyTimes()

	sub1 := NewMockSubscriber(ctrl)
	sub1called := false
	sub1.EXPECT().Dispatch(event1).Do(func(event eventBus.Event) {
		sub1called = true
	}).Return(nil)

	sub2 := NewMockSubscriber(ctrl)
	sub2.EXPECT().Dispatch(event2).
		Times(0) // publishing event1 won't trigger sub2

	bus := NewSimpleEventBus()

	bus.Subscribe(sub1, event1)
	bus.Subscribe(sub2, event2)

	bus.Publish(event1)
	bus.Wait()

	assert.True(t, sub1called)
}
