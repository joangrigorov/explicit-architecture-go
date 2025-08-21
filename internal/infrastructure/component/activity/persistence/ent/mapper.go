package ent

import (
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/component/activity/persistence/ent/generated"
)

func mapEntity(dto *generated.Activity) *domain.Activity {
	return domain.ReconstituteActivity(
		domain.ActivityID(dto.ID.String()),
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
