package activities

import (
	"app/internal/core/component/activity/domain"
	. "app/internal/presentation/api/component/activity/v1/responses"
	"app/internal/presentation/api/shared/responses"
	"app/mock/core/component/activity/application/repositories"
	"app/mock/infrastructure/framework/http"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestController_GetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	goCtx := context.Background()

	activityRepository := repositories.NewMockActivityRepository(ctrl)
	ctx := http.NewMockContext(ctrl)
	controller := Controller{activityRepository: activityRepository}

	f := faker.New()

	t.Run("not found error", func(t *testing.T) {
		ctx.EXPECT().ParamString("id").Return(f.UUID().V4())
		ctx.EXPECT().Context().Return(goCtx)
		activityRepository.EXPECT().
			GetById(goCtx, gomock.AssignableToTypeOf(domain.ActivityID(""))).
			Return(nil, errors.New("not found"))

		ctx.EXPECT().JSON(404, responses.DefaultError{Error: "not found"})

		controller.GetOne(ctx)
	})

	t.Run("success", func(t *testing.T) {
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
			ID:               domain.ActivityID(id),
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

		ctx.EXPECT().ParamString("id").Return(id)
		ctx.EXPECT().Context().Return(goCtx)
		activityRepository.EXPECT().
			GetById(goCtx, gomock.AssignableToTypeOf(domain.ActivityID(""))).
			Return(activity, nil)

		ctx.EXPECT().
			JSON(200, gomock.AssignableToTypeOf(&ActivityResponse{})).
			Do(func(statusCode int, resp *ActivityResponse) {
				assert.Equal(t, slug, resp.Slug)
				assert.Equal(t, title, resp.Title)
				assert.Equal(t, posterImageUrl, resp.PosterImageUrl)
				assert.Equal(t, shortDescription, resp.ShortDescription)
				assert.Equal(t, fullDescription, resp.FullDescription)
				assert.Equal(t, happensAt, resp.HappensAt)
				assert.Equal(t, attendants, resp.Attendants)
				assert.Equal(t, createdAt, resp.CreatedAt)
				assert.Equal(t, updatedAt, resp.UpdatedAt)
			})

		controller.GetOne(ctx)
	})
}
