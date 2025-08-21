package ent

import (
	"app/internal/core/component/attendance/domain"
	"app/internal/infrastructure/component/attendance/persistence/ent/generated"
	"app/internal/infrastructure/framework/uuid"
)

func mapEntity(dto *generated.Attendance) *domain.Attendance {
	return domain.ReconstituteAttendance(
		domain.AttendanceID(dto.ID.String()),
		domain.AttendeeID(dto.AttendeeID.String()),
		domain.ActivityID(dto.ActivityID.String()),
		dto.ActivitySlug,
		dto.ActivityTitle,
		dto.ActivityPosterImageURL,
		dto.ActivityShortDescription,
		dto.ActivityHappensAt,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}

func mapDto(at *domain.Attendance) *generated.Attendance {
	return &generated.Attendance{
		ID:                       uuid.Parse(at.ID),
		AttendeeID:               uuid.Parse(at.Attendee.ID),
		ActivityID:               uuid.Parse(at.Activity.ID),
		ActivitySlug:             at.Activity.Slug,
		ActivityTitle:            at.Activity.Title,
		ActivityPosterImageURL:   at.Activity.PosterImageUrl,
		ActivityShortDescription: at.Activity.ShortDescription,
		ActivityHappensAt:        at.Activity.HappensAt,
		CreatedAt:                at.CreatedAt,
		UpdatedAt:                at.UpdatedAt,
	}
}
