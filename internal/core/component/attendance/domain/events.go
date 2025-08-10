package domain

import "time"

type AttendanceEvent interface {
	Name() string
	CreatedAt() time.Time
}

type AttendanceCreated struct {
	createdAt time.Time
}

func NewAttendanceCreated() *AttendanceCreated {
	return &AttendanceCreated{createdAt: time.Now()}
}

func (a *AttendanceCreated) Name() string {
	return "Attendance.AttendanceCreated"
}

func (a *AttendanceCreated) CreatedAt() time.Time {
	return a.createdAt
}
