package uuid

import (
	port "app/internal/core/port/uuid"

	"github.com/google/uuid"
)

type Generator struct{}

func NewGenerator() port.Generator {
	return &Generator{}
}

func (g *Generator) Generate() string {
	return uuid.New().String()
}
