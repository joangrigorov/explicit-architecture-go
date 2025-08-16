package activity

import (
	"app/internal/core/component/activity/domain"
	"app/internal/infrastructure/persistence/ent/generated/activity"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestMapEntity(t *testing.T) {
	f := faker.New()

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
	deletedAt := f.Time().Time(happensAt)

	ac := mapEntity(&activity.Activity{
		ID:               id,
		Slug:             slug,
		Title:            title,
		PosterImageURL:   posterImageUrl,
		ShortDescription: shortDescription,
		FullDescription:  fullDescription,
		HappensAt:        happensAt,
		Attendants:       attendants,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        &deletedAt,
	})

	assert.NotNil(t, ac)

	assert.Equal(t, domain.ActivityID(id.String()), ac.ID)
	assert.Equal(t, slug, ac.Slug)
	assert.Equal(t, title, ac.Title)
	assert.Equal(t, posterImageUrl, ac.PosterImageUrl)
	assert.Equal(t, shortDescription, ac.ShortDescription)
	assert.Equal(t, fullDescription, ac.FullDescription)
	assert.Equal(t, happensAt, ac.HappensAt)
	assert.Equal(t, attendants, ac.Attendants)
	assert.Equal(t, createdAt, ac.CreatedAt)
	assert.Equal(t, updatedAt, ac.UpdatedAt)
}
