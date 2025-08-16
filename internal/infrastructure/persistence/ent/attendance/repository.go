package attendance

import (
	"app/internal/core/component/attendance/domain"
	"app/internal/infrastructure/framework/uuid"
	"app/internal/infrastructure/persistence/ent/generated/attendance"
	ent "app/internal/infrastructure/persistence/ent/generated/attendance/attendance"
	"context"
	"time"
)

type Repository struct {
	client *attendance.Client
}

func NewAttendanceRepository(client *attendance.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) GetById(ctx context.Context, id domain.AttendanceID) (*domain.Attendance, error) {
	dto, err := r.client.Attendance.
		Query().
		Where(
			ent.ID(uuid.Parse(id)),
			ent.DeletedAtIsNil(),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return mapEntity(dto), nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Attendance, error) {
	entries, err := r.client.Attendance.
		Query().
		Where(ent.DeletedAtIsNil()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	collection := make([]*domain.Attendance, 0, len(entries))

	for _, e := range entries {
		collection = append(collection, mapEntity(e))
	}

	return collection, nil
}

func (r *Repository) Create(ctx context.Context, at *domain.Attendance) error {
	builder := r.client.Attendance.Create()

	_, err := builder.
		SetID(uuid.Parse(at.ID)).
		SetAttendeeID(uuid.Parse(at.Attendee.ID)).
		SetActivityID(uuid.Parse(at.Activity.ID)).
		SetActivitySlug(at.Activity.Slug).
		SetActivityTitle(at.Activity.Title).
		SetActivityShortDescription(at.Activity.ShortDescription).
		SetActivityPosterImageURL(at.Activity.PosterImageUrl).
		SetActivityHappensAt(at.Activity.HappensAt).
		SetCreatedAt(at.CreatedAt).
		SetUpdatedAt(at.UpdatedAt).
		Save(ctx)

	return err
}

func (r *Repository) Update(ctx context.Context, at *domain.Attendance) error {
	dto := mapDto(at)

	_, err := r.client.Attendance.UpdateOne(dto).Save(ctx)

	return err
}

func (r *Repository) Delete(ctx context.Context, at *domain.Attendance) error {
	_, err := r.client.Attendance.
		UpdateOneID(uuid.Parse(at.ID)).
		SetDeletedAt(time.Now()).
		Save(ctx)

	return err
}
