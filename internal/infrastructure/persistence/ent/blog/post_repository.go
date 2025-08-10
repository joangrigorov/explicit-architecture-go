package blog

import (
	"app/internal/core/component/blog/application/repositories"
	"app/internal/core/component/blog/domain/post"
	"app/internal/infrastructure/persistence/ent/generated/blog"
	schema "app/internal/infrastructure/persistence/ent/generated/blog/post"
	"context"
)

type PostRepository struct {
	client *blog.Client
}

func NewPostRepository(client *blog.Client) repositories.PostRepository {
	return &PostRepository{client: client}
}

func (p PostRepository) GetById(ctx context.Context, id int) (*post.Post, error) {
	dto, err := p.client.Post.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return post.Map(dto.ID, dto.Name, dto.Content), err
}

func (p PostRepository) Create(ctx context.Context, post *post.Post) error {
	_, err := p.client.Post.
		Create().
		SetName(post.Name).
		SetContent(post.Content).
		Save(ctx)
	return err
}

func (p PostRepository) Update(ctx context.Context, post *post.Post) error {
	_, err := p.client.Post.
		Update().
		Where(schema.IDEQ(post.Id)).
		SetName(post.Name).
		SetContent(post.Content).
		Save(ctx)

	return err
}

func (p PostRepository) GetAll(ctx context.Context) ([]*post.Post, error) {
	entries, err := p.client.Post.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	responses := make([]*post.Post, 0, len(entries))
	for _, dto := range entries {
		responses = append(responses, post.Map(dto.ID, dto.Name, dto.Content))
	}

	return responses, nil
}

func (p PostRepository) Delete(ctx context.Context, post *post.Post) error {
	return p.client.Post.DeleteOneID(post.Id).Exec(ctx)
}
