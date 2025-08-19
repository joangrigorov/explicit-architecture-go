package uuid

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	generator := &Generator{}
	id := generator.Generate()

	_, err := uuid.Parse(id)

	assert.NoError(t, err)
}
