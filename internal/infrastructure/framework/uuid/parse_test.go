package uuid

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	type ExampleId string

	f := faker.New()
	t.Run("success", func(t *testing.T) {
		id := f.UUID().V4()
		uuid := Parse(id)

		assert.Equal(t, id, uuid.String())
	})

	t.Run("success (wrapped)", func(t *testing.T) {
		id := f.UUID().V4()
		uuid := Parse(ExampleId(id))

		assert.Equal(t, id, uuid.String())
	})

	t.Run("panics on bad format", func(t *testing.T) {
		id := "not an uuid"

		assert.Panics(t, func() {
			Parse(id)
		})
	})
}
