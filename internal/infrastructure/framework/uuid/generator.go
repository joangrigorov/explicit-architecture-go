package uuid

import (
	"github.com/google/uuid"
)

type Generator struct{}

func (g *Generator) Generate() string {
	return uuid.UUID{}.String()
}
