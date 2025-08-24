package hmac

import (
	port "app/internal/core/port/hmac"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Generator struct{}

func NewGenerator() port.Generator {
	return &Generator{}
}

func (g *Generator) Generate(message string) (string, string, error) {
	secret := make([]byte, 30)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))

	return hex.EncodeToString(h.Sum(nil)), hex.EncodeToString(secret), nil
}
