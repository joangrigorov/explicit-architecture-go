package domain

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var f = faker.New()

func TestNewActivity(t *testing.T) {
	id := ActivityID(f.UUID().V4())
	slug := f.Lorem().Word() + "-" + f.Lorem().Word()
	title := f.Lorem().Sentence(3)
	posterImageUrl := f.ProfileImage().Image().Name()
	shortDescription := f.Lorem().Sentence(10)
	fullDescription := f.Lorem().Sentence(50)
	happensAt := f.Time().Time(time.Now())
	attendants := f.Int()

	activity := NewActivity(
		id,
		slug,
		title,
		posterImageUrl,
		shortDescription,
		fullDescription,
		happensAt,
		attendants,
	)

	assert.NotNil(t, activity)

	assert.Equal(t, id, activity.ID)
	assert.Equal(t, slug, activity.Slug)
	assert.Equal(t, title, activity.Title)
	assert.Equal(t, posterImageUrl, activity.PosterImageUrl)
	assert.Equal(t, shortDescription, activity.ShortDescription)
	assert.Equal(t, fullDescription, activity.FullDescription)
	assert.Equal(t, happensAt, activity.HappensAt)
	assert.Equal(t, attendants, activity.Attendants)
	assert.NotNil(t, activity.CreatedAt)
	assert.NotNil(t, activity.UpdatedAt)
}

func TestReconstituteActivity(t *testing.T) {
	id := ActivityID(f.UUID().V4())
	slug := f.Lorem().Word() + "-" + f.Lorem().Word()
	title := f.Lorem().Sentence(3)
	posterImageUrl := f.ProfileImage().Image().Name()
	shortDescription := f.Lorem().Sentence(10)
	fullDescription := f.Lorem().Sentence(50)
	happensAt := f.Time().Time(time.Now())
	attendants := f.Int()
	createdAt := f.Time().Time(happensAt)
	updatedAt := f.Time().Time(happensAt)

	activity := ReconstituteActivity(
		id,
		slug,
		title,
		posterImageUrl,
		shortDescription,
		fullDescription,
		happensAt,
		attendants,
		createdAt,
		updatedAt,
	)

	assert.NotNil(t, activity)

	assert.Equal(t, id, activity.ID)
	assert.Equal(t, slug, activity.Slug)
	assert.Equal(t, title, activity.Title)
	assert.Equal(t, posterImageUrl, activity.PosterImageUrl)
	assert.Equal(t, shortDescription, activity.ShortDescription)
	assert.Equal(t, fullDescription, activity.FullDescription)
	assert.Equal(t, happensAt, activity.HappensAt)
	assert.Equal(t, attendants, activity.Attendants)
	assert.Equal(t, createdAt, activity.CreatedAt)
	assert.Equal(t, updatedAt, activity.UpdatedAt)
}
