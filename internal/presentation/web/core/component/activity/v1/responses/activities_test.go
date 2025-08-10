package responses

import (
	"app/internal/core/component/activity/domain"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	f := faker.New()

	id := f.UUID().V4()
	slug := f.Lorem().Word() + "-" + f.Lorem().Word()
	title := f.Lorem().Sentence(3)
	posterImageUrl := f.ProfileImage().Image().Name()
	shortDescription := f.Lorem().Sentence(10)
	fullDescription := f.Lorem().Sentence(50)
	happensAt := f.Time().Time(time.Now())
	attendants := f.Int()
	createdAt := f.Time().Time(time.Now())
	updatedAt := f.Time().Time(time.Now())

	activity := &domain.Activity{
		Id:               domain.ActivityId(id),
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

	response := One(activity)

	assert.NotNil(t, response)
	assert.Equal(t, id, response.Id)
	assert.Equal(t, slug, response.Slug)
	assert.Equal(t, title, response.Title)
	assert.Equal(t, posterImageUrl, response.PosterImageUrl)
	assert.Equal(t, shortDescription, response.ShortDescription)
	assert.Equal(t, fullDescription, response.FullDescription)
	assert.Equal(t, happensAt, response.HappensAt)
	assert.Equal(t, attendants, response.Attendants)
	assert.Equal(t, createdAt, response.CreatedAt)
	assert.Equal(t, updatedAt, response.UpdatedAt)
}

func TestMany(t *testing.T) {
	f := faker.New()

	id := f.UUID().V4()
	slug := f.Lorem().Word() + "-" + f.Lorem().Word()
	title := f.Lorem().Sentence(3)
	posterImageUrl := f.ProfileImage().Image().Name()
	shortDescription := f.Lorem().Sentence(10)
	fullDescription := f.Lorem().Sentence(50)
	happensAt := f.Time().Time(time.Now())
	attendants := f.Int()
	createdAt := f.Time().Time(time.Now())
	updatedAt := f.Time().Time(time.Now())

	activities := []*domain.Activity{{
		Id:               domain.ActivityId(id),
		Slug:             slug,
		Title:            title,
		PosterImageUrl:   posterImageUrl,
		ShortDescription: shortDescription,
		FullDescription:  fullDescription,
		HappensAt:        happensAt,
		Attendants:       attendants,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}}

	response := Many(activities)

	assert.NotNil(t, response)
	assert.Len(t, response, 1)

	assert.Equal(t, id, response[0].Id)
	assert.Equal(t, slug, response[0].Slug)
	assert.Equal(t, title, response[0].Title)
	assert.Equal(t, posterImageUrl, response[0].PosterImageUrl)
	assert.Equal(t, shortDescription, response[0].ShortDescription)
	assert.Equal(t, fullDescription, response[0].FullDescription)
	assert.Equal(t, happensAt, response[0].HappensAt)
	assert.Equal(t, attendants, response[0].Attendants)
	assert.Equal(t, createdAt, response[0].CreatedAt)
	assert.Equal(t, updatedAt, response[0].UpdatedAt)
}
