package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoles(t *testing.T) {
	assert.Equal(t, "member", Member{}.ID().String())
	assert.Equal(t, "admin", Admin{}.ID().String())
}
