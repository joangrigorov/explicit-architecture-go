package events

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jaswdr/faker/v2"
)

func TestEventID_String(t *testing.T) {
	f := faker.New()

	id := f.UUID().V4()
	e := EventID(id)

	assert.Equal(t, id, e.String())
}

type exampleEvent struct{}

func (e exampleEvent) ID() EventID {
	panic("irrelevant for the test")
}

func TestMakeEventID(t *testing.T) {
	assert.Equal(t, "app/internal/core/shared_kernel/events.exampleEvent", makeEventID(&exampleEvent{}).String())
	assert.Equal(t, "app/internal/core/shared_kernel/events.exampleEvent", makeEventID(exampleEvent{}).String())
}
