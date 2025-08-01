package requests

import "app/internal/core/component/blog/domain/post"

type CreatePost struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (r *CreatePost) ToPost() *post.Post {
	return &post.Post{
		Name:    r.Name,
		Content: r.Content,
	}
}
