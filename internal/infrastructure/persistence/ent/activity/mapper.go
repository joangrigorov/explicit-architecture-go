package activity

import (
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/framework/uuid"
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
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}

func mapDTO(ac *domain.Activity) *activity.Activity {
	return &activity.Activity{
		ID:               uuid.Parse(ac.Id),
		Slug:             ac.Slug,
		Title:            ac.Title,
		PosterImageURL:   ac.PosterImageUrl,
		ShortDescription: ac.ShortDescription,
		FullDescription:  ac.FullDescription,
		HappensAt:        ac.HappensAt,
		Attendants:       ac.Attendants,
		CreatedAt:        ac.CreatedAt,
		UpdatedAt:        ac.UpdatedAt,
	}
}
