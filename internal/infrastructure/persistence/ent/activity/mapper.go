package activity

import (
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/persistence/ent/generated/activity"
)

func mapEntity(dto *activity.Activity) *domain.Activity {
	return domain.ReconstituteActivity(
		domain.ActivityId(dto.ID.String()),
		dto.Slug,
		dto.Title,
		dto.PosterImageURL,
		dto.ShortDescription,
		dto.FullDescription,
		dto.HappensAt,
		dto.Attendants,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}
