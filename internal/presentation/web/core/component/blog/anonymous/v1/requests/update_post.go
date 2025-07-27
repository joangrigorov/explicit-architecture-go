package requests

import (
	"app/internal/core/component/blog/domain/post"
)

type UpdatePostRequest struct {
	Name    *string `json:"name"`
	Content *string `json:"content"`
}

func (r *UpdatePostRequest) Populate(p *post.Post) {
	if r.Name != nil {
		p.Name = *r.Name
	}
	if r.Content != nil {
		p.Content = *r.Content
	}
}
