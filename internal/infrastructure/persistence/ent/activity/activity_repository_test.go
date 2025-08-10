package activity

import (
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/persistence/ent/generated/activity"
	"app/internal/infrastructure/persistence/ent/generated/activity/enttest"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewActivityRepository(t *testing.T) {
	r := NewActivityRepository(&activity.Client{})

	assert.NotNil(t, r)
}

func TestActivityRepository_GetById(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	t.Run("exists", func(t *testing.T) {
		id, slug, title, posterImageUrl, shortDescription, fullDescription, happensAt, attendants, createdAt, updatedAt := fakeActivityData(f)

		_, err := client.Activity.Create().
			SetID(id).
			SetSlug(slug).
			SetTitle(title).
			SetPosterImageURL(posterImageUrl).
			SetShortDescription(shortDescription).
			SetFullDescription(fullDescription).
			SetAttendants(attendants).
			SetCreatedAt(createdAt).
			SetUpdatedAt(updatedAt).
			SetHappensAt(happensAt).
			Save(ctx)
		if err != nil {
			t.Fatal(err)
		}

		repo := &ActivityRepository{client: client}

		got, err := repo.GetById(ctx, domain.ActivityId(id.String()))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, domain.ActivityId(id.String()), got.Id)
		assert.Equal(t, slug, got.Slug)
		assert.Equal(t, title, got.Title)
		assert.Equal(t, posterImageUrl, got.PosterImageUrl)
		assert.Equal(t, shortDescription, got.ShortDescription)
		assert.Equal(t, fullDescription, got.FullDescription)
		assert.Equal(t, happensAt, got.HappensAt.In(time.Local))
		assert.Equal(t, attendants, got.Attendants)
		assert.Equal(t, createdAt, got.CreatedAt.In(time.Local))
		assert.Equal(t, updatedAt, got.UpdatedAt.In(time.Local))
	})

	t.Run("soft deleted", func(t *testing.T) {
		_, err := client.Activity.Create().
			SetID(uuid.New()).
			SetSlug(f.Lorem().Word() + "-" + f.Lorem().Word()).
			SetTitle(f.Lorem().Sentence(3)).
			SetPosterImageURL(f.ProfileImage().Image().Name()).
			SetShortDescription(f.Lorem().Sentence(10)).
			SetFullDescription(f.Lorem().Sentence(50)).
			SetAttendants(f.Int()).
			SetCreatedAt(f.Time().Time(f.Time().Time(time.Now()))).
			SetUpdatedAt(f.Time().Time(f.Time().Time(time.Now()))).
			SetDeletedAt(f.Time().Time(f.Time().Time(time.Now()))).
			SetHappensAt(f.Time().Time(time.Now())).
			Save(ctx)
		if err != nil {
			t.Fatal(err)
		}

		repo := &ActivityRepository{client: client}

		_, err = repo.GetById(ctx, domain.ActivityId(uuid.New().String()))

		assert.ErrorContains(t, err, "activity not found")
	})
}

func TestActivityRepository_Create(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	id, slug, title, posterImageUrl, shortDescription, fullDescription, happensAt, attendants, createdAt, updatedAt := fakeActivityData(f)

	ac := &domain.Activity{
		Id:               domain.ActivityId(id.String()),
		Slug:             slug,
		Title:            title,
		PosterImageUrl:   posterImageUrl,
		ShortDescription: shortDescription,
		FullDescription:  fullDescription,
		HappensAt:        happensAt,
		Attendants:       attendants,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}

	r := &ActivityRepository{client: client}

	assert.NoError(t, r.Create(ctx, ac))
}

func TestActivityRepository_Update(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	id := uuid.New()

	_, err := client.Activity.Create().
		SetID(id).
		SetSlug(f.Lorem().Word() + "-" + f.Lorem().Word()).
		SetTitle(f.Lorem().Sentence(3)).
		SetPosterImageURL(f.ProfileImage().Image().Name()).
		SetShortDescription(f.Lorem().Sentence(10)).
		SetFullDescription(f.Lorem().Sentence(50)).
		SetAttendants(f.Int()).
		SetCreatedAt(f.Time().Time(f.Time().Time(time.Now()))).
		SetUpdatedAt(f.Time().Time(f.Time().Time(time.Now()))).
		SetHappensAt(f.Time().Time(time.Now())).
		Save(ctx)

	assert.NoError(t, err)

	_, slug, title, posterImageUrl, shortDescription, fullDescription, happensAt, attendants, createdAt, _ := fakeActivityData(f)

	ac := &domain.Activity{
		Id:               domain.ActivityId(id.String()),
		Slug:             slug,
		Title:            title,
		PosterImageUrl:   posterImageUrl,
		ShortDescription: shortDescription,
		FullDescription:  fullDescription,
		HappensAt:        happensAt,
		Attendants:       attendants,
		CreatedAt:        createdAt,
	}

	r := &ActivityRepository{client: client}
	assert.NoError(t, r.Update(ctx, ac))

	assert.NotNil(t, ac.UpdatedAt)

	dto, err := client.Activity.Get(ctx, id)
	assert.NoError(t, err)

	assert.Equal(t, slug, dto.Slug)
	assert.Equal(t, title, dto.Title)
	assert.Equal(t, posterImageUrl, dto.PosterImageURL)
	assert.Equal(t, shortDescription, dto.ShortDescription)
	assert.Equal(t, fullDescription, dto.FullDescription)
	assert.Equal(t, happensAt, dto.HappensAt.In(time.Local))
	assert.Equal(t, attendants, dto.Attendants)
	assert.Equal(t, createdAt, dto.CreatedAt.In(time.Local))
	assert.Nil(t, dto.DeletedAt)
}

func TestActivityRepository_GetAll(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	id, slug, title, posterImageUrl, shortDescription, fullDescription, happensAt, attendants, createdAt, updatedAt := fakeActivityData(f)

	// used to check data integrity
	_, err := client.Activity.Create().
		SetID(id).
		SetSlug(slug).
		SetTitle(title).
		SetPosterImageURL(posterImageUrl).
		SetShortDescription(shortDescription).
		SetFullDescription(fullDescription).
		SetAttendants(attendants).
		SetCreatedAt(createdAt).
		SetUpdatedAt(updatedAt).
		SetHappensAt(happensAt).
		Save(ctx)

	assert.NoError(t, err)

	// fillter
	for i := 0; i < 5; i++ {
		_, err := client.Activity.Create().
			SetID(uuid.New()).
			SetSlug(f.Lorem().Word() + "-" + f.Lorem().Word()).
			SetTitle(f.Lorem().Sentence(3)).
			SetPosterImageURL(f.ProfileImage().Image().Name()).
			SetShortDescription(f.Lorem().Sentence(10)).
			SetFullDescription(f.Lorem().Sentence(50)).
			SetAttendants(f.Int()).
			SetCreatedAt(f.Time().Time(f.Time().Time(time.Now()))).
			SetUpdatedAt(f.Time().Time(f.Time().Time(time.Now()))).
			SetHappensAt(f.Time().Time(time.Now())).
			Save(ctx)

		assert.NoError(t, err)
	}

	// soft-deleted:L shouldn't be included
	_, err = client.Activity.Create().
		SetID(uuid.New()).
		SetSlug(f.Lorem().Word() + "-" + f.Lorem().Word()).
		SetTitle(f.Lorem().Sentence(3)).
		SetPosterImageURL(f.ProfileImage().Image().Name()).
		SetShortDescription(f.Lorem().Sentence(10)).
		SetFullDescription(f.Lorem().Sentence(50)).
		SetAttendants(f.Int()).
		SetCreatedAt(f.Time().Time(f.Time().Time(time.Now()))).
		SetUpdatedAt(f.Time().Time(f.Time().Time(time.Now()))).
		SetHappensAt(f.Time().Time(time.Now())).
		SetDeletedAt(time.Now()).
		Save(ctx)

	assert.NoError(t, err)

	r := &ActivityRepository{client: client}

	entries, err := r.GetAll(ctx)

	assert.NoError(t, err)

	assert.Len(t, entries, 6)
	assert.Equal(t, id.String(), string(entries[0].Id))
	assert.Equal(t, slug, entries[0].Slug)
	assert.Equal(t, title, entries[0].Title)
	assert.Equal(t, posterImageUrl, entries[0].PosterImageUrl)
	assert.Equal(t, shortDescription, entries[0].ShortDescription)
	assert.Equal(t, fullDescription, entries[0].FullDescription)
	assert.Equal(t, happensAt, entries[0].HappensAt.In(time.Local))
	assert.Equal(t, attendants, entries[0].Attendants)
	assert.Equal(t, createdAt, entries[0].CreatedAt.In(time.Local))
	assert.Equal(t, updatedAt, entries[0].UpdatedAt.In(time.Local))
}

func TestActivityRepository_Delete(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", ":memory:?_fk=1")
	defer client.Close()

	f := faker.New()

	id := uuid.New()

	_, err := client.Activity.Create().
		SetID(id).
		SetSlug(f.Lorem().Word() + "-" + f.Lorem().Word()).
		SetTitle(f.Lorem().Sentence(3)).
		SetPosterImageURL(f.ProfileImage().Image().Name()).
		SetShortDescription(f.Lorem().Sentence(10)).
		SetFullDescription(f.Lorem().Sentence(50)).
		SetAttendants(f.Int()).
		SetCreatedAt(f.Time().Time(f.Time().Time(time.Now()))).
		SetUpdatedAt(f.Time().Time(f.Time().Time(time.Now()))).
		SetHappensAt(f.Time().Time(time.Now())).
		Save(ctx)

	assert.NoError(t, err)

	r := &ActivityRepository{client: client}
	err = r.Delete(ctx, &domain.Activity{Id: domain.ActivityId(id.String())})
	assert.NoError(t, err)

	dto, err := client.Activity.Get(ctx, id)

	assert.NoError(t, err)

	assert.NotNil(t, dto.DeletedAt)
}

func fakeActivityData(f faker.Faker) (
	uuid.UUID,
	string,
	string,
	string,
	string,
	string,
	time.Time,
	int,
	time.Time,
	time.Time,
) {
	id := uuid.New()
	slug := f.Lorem().Word() + "-" + f.Lorem().Word()
	title := f.Lorem().Sentence(3)
	posterImageUrl := f.ProfileImage().Image().Name()
	shortDescription := f.Lorem().Sentence(10)
	fullDescription := f.Lorem().Sentence(50)
	happensAt := f.Time().Time(time.Now())
	attendants := f.Int()
	createdAt := f.Time().Time(happensAt)
	updatedAt := f.Time().Time(happensAt)

	return id, slug, title, posterImageUrl, shortDescription, fullDescription, happensAt, attendants, createdAt, updatedAt
}
