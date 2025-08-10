package requests

import (
	"app/internal/core/component/activity/domain"
	"testing"

	optional "github.com/aarondl/null/v9"
	"github.com/stretchr/testify/assert"
)

func TestCreateUpdatePostRequest_Populate(t *testing.T) {
	p := &domain.Activity{}

	(&UpdatePostRequest{
		Name:    optional.NewString("My Title", true),
		Content: optional.NewString("My Content", true),
	}).Populate(p)

	assert.Equal(t, "My Title", p.Name)
	assert.Equal(t, "My Content", p.Content)

	(&UpdatePostRequest{
		Name: optional.NewString("New title", true),
	}).Populate(p)

	assert.Equal(t, "New title", p.Name)
	assert.Equal(t, "My Content", p.Content)

	(&UpdatePostRequest{
		Content: optional.NewString("New content", true),
	}).Populate(p)

	assert.Equal(t, "New title", p.Name)
	assert.Equal(t, "New content", p.Content)

	(&UpdatePostRequest{}).Populate(p)

	assert.Equal(t, "New title", p.Name)
	assert.Equal(t, "New content", p.Content)
}
