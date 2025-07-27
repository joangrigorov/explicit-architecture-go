package repositories

import (
	"app/internal/core/component/blog/domain/post"
	"context"
)

type PostRepository interface {
	GetById(ctx context.Context, id int) (*post.Post, error)
	Create(ctx context.Context, post *post.Post) error
}
