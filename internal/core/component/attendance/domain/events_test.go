package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttendanceCreated(t *testing.T) {
	event := NewAttendanceCreated()

	assert.NotNil(t, event)

	assert.Equal(t, "Attendance.AttendanceCreated", event.Name())
	assert.NotNil(t, event.CreatedAt())
}
