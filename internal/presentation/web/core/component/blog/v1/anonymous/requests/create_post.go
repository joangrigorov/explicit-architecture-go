package requests

import (
	"app/internal/core/component/activity/domain"
)

type CreatePost struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (r *CreatePost) ToPost() *domain.Activity {
	return &domain.Activity{
		Name:    r.Name,
		Content: r.Content,
	}
}
