package shared_kernel

import (
	"app/internal/presentation/web/core/component/blog/v1/anonymous/controllers/post"
	ctx "app/internal/presentation/web/port/http"
	"app/mock/core/component/blog/application/repositories"
	ut "app/mock/ext/go-playground/universal-translator"
	"app/mock/presentation/web/port/http"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	router := http.NewMockRouter(ctrl)
	group := http.NewMockRouter(ctrl)
	postRepository := repositories.NewMockPostRepository(ctrl)
	translator := ut.NewMockTranslator(ctrl)

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
		RegisterRoutes(router, post.NewController(postRepository, translator))
	})
}
