package repositories

import (
	"app/internal/core/component/attendance/domain"
	"context"
)

type AttendanceRepository interface {
	GetById(context.Context, domain.AttendanceID) (*domain.Attendance, error)
	GetAll(context.Context) ([]*domain.Attendance, error)
	Create(context.Context, *domain.Attendance) error
	Update(context.Context, *domain.Attendance) error
	Delete(context.Context, *domain.Attendance) error
}
