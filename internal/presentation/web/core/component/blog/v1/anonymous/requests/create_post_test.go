package requests

import (
	"app/internal/core/component/activity/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreatePostRequest_ToPost verifies that ToPost() maps fields correctly.
func TestCreatePostRequest_ToPost(t *testing.T) {
	req := &CreatePost{
		Name:    "My Title",
		Content: "My content here",
	}

	expected := &domain.Activity{
		Name:    "My Title",
		Content: "My content here",
	}

	actual := req.ToPost()

	assert.Equal(t, expected, actual)
}
