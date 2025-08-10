package activity

import (
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/persistence/ent/generated/activity"
	"context"
	"time"
)

type ActivityRepository struct {
	client *activity.Client
}

func NewActivityRepository(client *activity.Client) *ActivityRepository {
	return &ActivityRepository{client: client}
}

func (r *ActivityRepository) GetById(ctx context.Context, id domain.ActivityId) (*domain.Activity, error) {
	dto, err := r.client.Activity.Get(ctx, uuid.Parse(id))

	if err != nil {
		return nil, err
	}

	return mapEntity(dto), nil
}

func (r *ActivityRepository) Create(ctx context.Context, ac *domain.Activity) error {
	builder := r.client.Activity.Create()

	_, err := builder.
		SetID(uuid.Parse(ac.Id)).
		SetSlug(ac.Slug).
		SetTitle(ac.Title).
		SetPosterImageURL(ac.PosterImageUrl).
		SetShortDescription(ac.ShortDescription).
		SetFullDescription(ac.FullDescription).
		SetHappensAt(ac.HappensAt).
		SetCreatedAt(ac.CreatedAt).
		SetUpdatedAt(ac.UpdatedAt).
		SetAttendants(0).
		Save(ctx)

	return err
}

func (r *ActivityRepository) Update(ctx context.Context, ac *domain.Activity) error {
	_, err := r.client.Activity.
		UpdateOne(mapDTO(ac)).
		Save(ctx)

	return err
}

func (r *ActivityRepository) GetAll(ctx context.Context) ([]*domain.Activity, error) {
	entries, err := r.client.Activity.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	collection := make([]*domain.Activity, len(entries))

	for _, dto := range entries {
		collection = append(collection, mapEntity(dto))
	}

	return collection, nil
}

func (r *ActivityRepository) Delete(ctx context.Context, ac *domain.Activity) error {
	_, err := r.client.Activity.
		UpdateOneID(uuid.Parse(ac.Id)).
		SetDeletedAt(time.Now()).
		Save(ctx)
	return err
}
