package attendance

import (
	"app/internal/core/component/attendance/domain"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/persistence/ent/generated/attendance"
)

func mapEntity(dto *attendance.Attendance) *domain.Attendance {
	return domain.ReconstituteAttendance(
		domain.AttendanceId(dto.ID.String()),
		domain.AttendeeId(dto.AttendeeID.String()),
		domain.ActivityId(dto.ActivityID.String()),
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
		ID:                       uuid.Parse(at.Id),
		AttendeeID:               uuid.Parse(at.Attendee.Id),
		ActivityID:               uuid.Parse(at.Activity.Id),
		ActivitySlug:             at.Activity.Slug,
		ActivityTitle:            at.Activity.Title,
		ActivityPosterImageURL:   at.Activity.PosterImageUrl,
		ActivityShortDescription: at.Activity.ShortDescription,
		ActivityHappensAt:        at.Activity.HappensAt,
		CreatedAt:                at.CreatedAt,
		UpdatedAt:                at.UpdatedAt,
	}
}
