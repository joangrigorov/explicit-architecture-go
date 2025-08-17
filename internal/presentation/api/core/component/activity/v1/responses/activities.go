package responses

import (
	"app/internal/core/component/activity/domain"
	"time"
)

type ActivityResponse struct {
	Id               string    `json:"id"`
	Slug             string    `json:"slug"`
	Title            string    `json:"title"`
	PosterImageUrl   string    `json:"poster_image_url"`
	ShortDescription string    `json:"short_description"`
	FullDescription  string    `json:"full_description"`
	HappensAt        time.Time `json:"happens_at"`
	Attendants       int       `json:"attendants"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func One(ac *domain.Activity) *ActivityResponse {
	return &ActivityResponse{
		Id:               string(ac.ID),
		Slug:             ac.Slug,
		Title:            ac.Title,
		PosterImageUrl:   ac.PosterImageUrl,
		ShortDescription: ac.ShortDescription,
		FullDescription:  ac.FullDescription,
		HappensAt:        ac.HappensAt,
		Attendants:       ac.Attendants,
		CreatedAt:        ac.CreatedAt,
		UpdatedAt:        ac.UpdatedAt,
	}
}

func Many(entries []*domain.Activity) []*ActivityResponse {
	responses := make([]*ActivityResponse, 0, len(entries))

	for _, ac := range entries {
		responses = append(responses, One(ac))
	}
	return responses
}
