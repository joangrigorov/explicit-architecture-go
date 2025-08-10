package post

import (
	"go.uber.org/mock/gomock"
	"testing"
)

func TestController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//mockPostRepo := mockrepo.NewMockPostRepository(ctrl)

	// Create controller with mock repo
	//c := NewController(mockPostRepo)

	t.Run("success", func(t *testing.T) {

	})

	t.Run("bind json error", func(t *testing.T) {

	})

	t.Run("repository error", func(t *testing.T) {

	})
}
