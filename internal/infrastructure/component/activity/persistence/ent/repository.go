package ent

import (
	"app/internal/core/component/activity/application/repositories"
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/component/activity/persistence/ent/generated"
	ent "app/internal/infrastructure/component/activity/persistence/ent/generated/activity"
	"app/internal/infrastructure/framework/uuid"
	"context"
	"time"
)

type Repository struct {
	client *generated.Client
}

func NewActivityRepository(client *generated.Client) repositories.ActivityRepository {
	return &Repository{client: client}
}

func (r *Repository) GetById(ctx context.Context, id domain.ActivityID) (*domain.Activity, error) {
	parse := uuid.Parse(id)

	dto, err := r.client.Activity.
		Query().
		Where(
			ent.ID(parse),
			ent.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return mapEntity(dto), nil
}

func (r *Repository) Create(ctx context.Context, ac *domain.Activity) error {
	builder := r.client.Activity.Create()

	_, err := builder.
		SetID(uuid.Parse(ac.ID)).
		SetSlug(ac.Slug).
		SetTitle(ac.Title).
		SetPosterImageURL(ac.PosterImageUrl).
		SetShortDescription(ac.ShortDescription).
		SetFullDescription(ac.FullDescription).
		SetHappensAt(ac.HappensAt).
		SetCreatedAt(ac.CreatedAt).
		SetUpdatedAt(ac.UpdatedAt).
		SetAttendants(ac.Attendants).
		Save(ctx)

	return err
}

func (r *Repository) Update(ctx context.Context, ac *domain.Activity) error {
	updatedAt := time.Now()
	_, err := r.client.Activity.
		UpdateOneID(uuid.Parse(ac.ID)).
		SetSlug(ac.Slug).
		SetTitle(ac.Title).
		SetPosterImageURL(ac.PosterImageUrl).
		SetShortDescription(ac.ShortDescription).
		SetFullDescription(ac.FullDescription).
		SetHappensAt(ac.HappensAt).
		SetAttendants(ac.Attendants).
		SetUpdatedAt(updatedAt).
		SetCreatedAt(ac.CreatedAt).
		Save(ctx)

	ac.UpdatedAt = updatedAt

	return err
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Activity, error) {
	entries, err := r.client.Activity.
		Query().
		Where(ent.DeletedAtIsNil()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	collection := make([]*domain.Activity, 0, len(entries))

	for _, dto := range entries {
		collection = append(collection, mapEntity(dto))
	}

	return collection, nil
}

func (r *Repository) Delete(ctx context.Context, ac *domain.Activity) error {
	_, err := r.client.Activity.
		UpdateOneID(uuid.Parse(ac.ID)).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}
