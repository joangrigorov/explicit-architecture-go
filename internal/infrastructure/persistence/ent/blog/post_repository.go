package blog

import (
	"app/internal/core/component/blog/application/repositories"
	"app/internal/core/component/blog/domain/post"
	"app/internal/infrastructure/persistence/ent/generated/blog"
	"context"
)

type PostRepository struct {
	client *blog.Client
}

func NewPostRepository(client *blog.Client) repositories.PostRepository {
	return &PostRepository{client: client}
}

func (p PostRepository) GetById(ctx context.Context, id int) (*post.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostRepository) Create(ctx context.Context, post *post.Post) error {
	_, err := p.client.Post.
		Create().
		SetName(post.Name).
		SetContent(post.Content).
		Save(ctx)
	return err
}
