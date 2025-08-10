package middleware

import (
	"app/mock/presentation/web/port/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestValidateJSONBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := http.NewMockContext(ctrl)

	t.Run("skips validation if method is not supported", func(t *testing.T) {
		ctx.EXPECT().IsPost().Return(false)
		ctx.EXPECT().IsPut().Return(false)
		ctx.EXPECT().IsPatch().Return(false)

		ctx.EXPECT().Next()

		ValidateJSONBody(ctx)
	})

	t.Run("skips validation if request content type is not application/json", func(t *testing.T) {
		ctx.EXPECT().IsPost().Return(true)
		ctx.EXPECT().IsJsonRequest().Return(false)

		ctx.EXPECT().Next()

		ValidateJSONBody(ctx)
	})

	t.Run("aborts on invalid json body", func(t *testing.T) {
		ctx.EXPECT().IsPost().Return(true)
		ctx.EXPECT().IsJsonRequest().Return(true)
		ctx.EXPECT().IsJsonBodyValid().Return(false)

		ctx.EXPECT().AbortWithStatusJSON(400, gin.H{"error": "invalid JSON"})

		ValidateJSONBody(ctx)
	})

	t.Run("validation passes and handler proceeds", func(t *testing.T) {
		ctx.EXPECT().IsPost().Return(true)
		ctx.EXPECT().IsJsonRequest().Return(true)
		ctx.EXPECT().IsJsonBodyValid().Return(true)

		ctx.EXPECT().Next()

		ValidateJSONBody(ctx)
	})
}

func TestRequiresJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := http.NewMockContext(ctrl)

	t.Run("aborts if request content type isn't application/json", func(t *testing.T) {
		ctx.EXPECT().IsJsonRequest().Return(false)

		ctx.EXPECT().AbortWithStatusJSON(400, gin.H{"error": "This API only serves JSON requests"})

		RequiresJSON(ctx)
	})

	t.Run("passes if request content type is application/json", func(t *testing.T) {
		ctx.EXPECT().IsJsonRequest().Return(true)

		ctx.EXPECT().Next()

		RequiresJSON(ctx)
	})
}
