package pgsql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUrl(t *testing.T) {
	t.Run("sslmode disabled", func(t *testing.T) {
		assert.Equal(
			t,
			"postgres://user:password@localhost:5432/my_db?sslmode=disable",
			buildUrl(
				"user",
				"password",
				"localhost",
				"5432",
				"my_db",
				"disable",
			),
		)
	})

	t.Run("sslmode verify-full", func(t *testing.T) {
		assert.Equal(
			t,
			"postgres://user:password@localhost:5432/my_db?sslmode=verify-full",
			buildUrl(
				"user",
				"password",
				"localhost",
				"5432",
				"my_db",
				"verify-full",
			),
		)
	})

	t.Run("panic on wrong sslmode", func(t *testing.T) {
		assert.Panics(t, func() {
			buildUrl(
				"user",
				"password",
				"localhost",
				"5432",
				"my_db",
				"not_supported",
			)
		})
	})
}
