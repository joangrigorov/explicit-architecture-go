package responses

import (
	"app/mock/infrastructure/http"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type e struct {
	message string
}

func TestErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("default error", func(t *testing.T) {
		response := NewDefaultError(errors.New("test"))
		assert.Equal(t, DefaultError{"test"}, response)
	})

	t.Run("Unprocessable entity", func(t *testing.T) {
		ctx := http.NewMockContext(ctrl)
		ctx.EXPECT().JSON(422, map[string]string{
			"foo": "bar",
		})
		UnprocessableEntity(ctx, map[string]string{
			"foo": "bar",
		})
	})

	t.Run("Internal server error", func(t *testing.T) {
		ctx := http.NewMockContext(ctrl)
		ctx.EXPECT().JSON(500, map[string]string{
			"foo": "bar",
		})
		InternalServerError(ctx, map[string]string{
			"foo": "bar",
		})
	})

	t.Run("Bad request", func(t *testing.T) {
		ctx := http.NewMockContext(ctrl)
		ctx.EXPECT().JSON(400, map[string]string{
			"foo": "bar",
		})
		BadRequest(ctx, map[string]string{
			"foo": "bar",
		})
	})

	t.Run("Not found", func(t *testing.T) {
		ctx := http.NewMockContext(ctrl)
		ctx.EXPECT().JSON(404, map[string]string{
			"foo": "bar",
		})
		NotFound(ctx, map[string]string{
			"foo": "bar",
		})
	})
}
