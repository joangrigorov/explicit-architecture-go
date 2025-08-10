package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewActivityCreated(t *testing.T) {
	event := NewActivityCreated()

	assert.NotNil(t, event)

	assert.Equal(t, "Activity.ActivityCreated", event.Name())
	assert.NotNil(t, event.CreatedAt())
}
