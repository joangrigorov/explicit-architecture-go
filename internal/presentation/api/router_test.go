package api

import (
	ctx "app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/activity/v1/controllers/activities"
	"app/mock/infrastructure/framework/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegisterRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	router := http.NewMockRouter(ctrl)
	group := http.NewMockRouter(ctrl)

	router.EXPECT().
		Use(gomock.Any()).
		DoAndReturn(func(...ctx.Handler) {}).
		AnyTimes()
	router.EXPECT().Use(gomock.Any()).AnyTimes()
	router.EXPECT().Group(gomock.Any()).Return(group).AnyTimes()

	group.EXPECT().
		POST(gomock.AssignableToTypeOf(""), gomock.Any()).
		DoAndReturn(func(string, ...ctx.Handler) {}).
		AnyTimes()

	group.EXPECT().
		PUT(gomock.AssignableToTypeOf(""), gomock.Any()).
		DoAndReturn(func(string, ...ctx.Handler) {}).
		AnyTimes()

	group.EXPECT().
		PATCH(gomock.AssignableToTypeOf(""), gomock.Any()).
		DoAndReturn(func(string, ...ctx.Handler) {}).
		AnyTimes()

	group.EXPECT().
		DELETE(gomock.AssignableToTypeOf(""), gomock.Any()).
		DoAndReturn(func(string, ...ctx.Handler) {}).
		AnyTimes()

	group.EXPECT().
		GET(gomock.AssignableToTypeOf(""), gomock.Any()).
		DoAndReturn(func(string, ...ctx.Handler) {}).
		AnyTimes()

	assert.NotPanics(t, func() {
		RegisterRoutes(
			router,
			&activities.Controller{},
		)
	})
}
