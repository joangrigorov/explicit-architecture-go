package verification

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

type Token []byte

func (t Token) Bytes() []byte {
	return t
}

func (t Token) Encode() string {
	return base64.RawURLEncoding.EncodeToString(t)
}

func (t Token) Hash() CSRFToken {
	return sha256.Sum256(t)
}

func GenerateToken() (Token, error) {
	const length = 32
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

type CSRFToken [32]byte

func (t CSRFToken) Encode() string {
	return base64.RawURLEncoding.EncodeToString(t[:])
}

func (t CSRFToken) Bytes() [32]byte {
	return t
}

func DecodeCSRFToken(s string) (CSRFToken, error) {
	var t CSRFToken
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return t, err
	}
	if len(b) != len(t) {
		return t, InvalidCSRFTokenLengthError{}
	}
	copy(t[:], b)
	return t, nil
}

type InvalidCSRFTokenLengthError struct{}

func (i InvalidCSRFTokenLengthError) Error() string {
	return "invalid CSRF token length"
}
