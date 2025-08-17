package activities

import (
	"app/internal/core/component/activity/domain"
	. "app/internal/presentation/api/core/component/activity/v1/responses"
	"app/internal/presentation/api/core/shared/responses"
	"app/mock/core/component/activity/application/repositories"
	"app/mock/presentation/api/port/http"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestController_Index(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	goCtx := context.Background()

	activityRepository := repositories.NewMockActivityRepository(ctrl)
	ctx := http.NewMockContext(ctrl)
	controller := Controller{activityRepository: activityRepository}

	f := faker.New()

	t.Run("internal server error", func(t *testing.T) {
		ctx.EXPECT().Context().Return(goCtx)
		activityRepository.EXPECT().
			GetAll(goCtx).
			Return(nil, errors.New("something went wrong"))
		ctx.EXPECT().JSON(500, &responses.DefaultError{Error: "Application error."})

		controller.Index(ctx)
	})

	t.Run("success (empty list)", func(t *testing.T) {
		ctx.EXPECT().Context().Return(goCtx)
		activityRepository.EXPECT().
			GetAll(goCtx).
			Return([]*domain.Activity{}, nil)

		ctx.EXPECT().JSON(200, make([]*ActivityResponse, 0))

		controller.Index(ctx)
	})

	t.Run("success (not empty)", func(t *testing.T) {
		ctx.EXPECT().Context().Return(goCtx)

		firstId := f.UUID().V4()
		slug := f.Lorem().Word() + "-" + f.Lorem().Word()
		title := f.Lorem().Sentence(3)
		posterImageUrl := f.ProfileImage().Image().Name()
		shortDescription := f.Lorem().Sentence(10)
		fullDescription := f.Lorem().Sentence(50)
		happensAt := f.Time().Time(time.Now())
		attendants := f.Int()
		createdAt := f.Time().Time(time.Now())
		updatedAt := f.Time().Time(time.Now())

		entries := []*domain.Activity{
			{
				ID:               domain.ActivityID(firstId),
				Slug:             slug,
				Title:            title,
				PosterImageUrl:   posterImageUrl,
				ShortDescription: shortDescription,
				FullDescription:  fullDescription,
				HappensAt:        happensAt,
				Attendants:       attendants,
				CreatedAt:        createdAt,
				UpdatedAt:        updatedAt,
			},
		}

		activityRepository.EXPECT().
			GetAll(goCtx).
			Return(entries, nil)

		ctx.EXPECT().
			JSON(200, gomock.AssignableToTypeOf([]*ActivityResponse{})).
			Do(func(statusCode int, entries []*ActivityResponse) {
				assert.Len(t, entries, 1)
				first := entries[0]

				assert.Equal(t, firstId, first.Id)
				assert.Equal(t, slug, first.Slug)
				assert.Equal(t, title, first.Title)
				assert.Equal(t, posterImageUrl, first.PosterImageUrl)
				assert.Equal(t, shortDescription, first.ShortDescription)
				assert.Equal(t, fullDescription, first.FullDescription)
				assert.Equal(t, happensAt, first.HappensAt)
				assert.Equal(t, attendants, first.Attendants)
				assert.Equal(t, createdAt, first.CreatedAt)
				assert.Equal(t, updatedAt, first.UpdatedAt)
			})

		controller.Index(ctx)
	})
}
