package attendance

import (
	"app/internal/core/component/attendance/domain"
	"app/internal/infrastructure/persistence/ent/generated/attendance"
	"app/internal/infrastructure/uuid"
)

func mapEntity(dto *attendance.Attendance) *domain.Attendance {
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

func mapDto(at *domain.Attendance) *attendance.Attendance {
	return &attendance.Attendance{
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
