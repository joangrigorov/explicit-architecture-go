package responses

import (
	"app/internal/core/component/activity/domain"
)

type PostResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func OnePostResponse(p *domain.Activity) *PostResponse {
	return &PostResponse{
		Id:      p.Id,
		Name:    p.Name,
		Content: p.Content,
	}
}

func MultiPostResponse(posts []*domain.Activity) []*PostResponse {
	responses := make([]*PostResponse, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, OnePostResponse(p))
	}
	return responses
}
