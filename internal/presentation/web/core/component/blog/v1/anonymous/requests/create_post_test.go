package requests

import (
	"app/internal/core/component/blog/domain/post"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCreatePostRequest_ToPost verifies that ToPost() maps fields correctly.
func TestCreatePostRequest_ToPost(t *testing.T) {
	req := &CreatePostRequest{
		Name:    "My Title",
		Content: "My content here",
	}

	expected := &post.Post{
		Name:    "My Title",
		Content: "My content here",
	}

	actual := req.ToPost()

	assert.Equal(t, expected, actual)
}
