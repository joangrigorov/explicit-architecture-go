package get_verification_preflight

import "time"

type VerificationPreflightDto struct {
	Valid       bool      `json:"valid"`
	MaskedEmail string    `json:"masked_email"`
	ExpiresAt   time.Time `json:"expires_at"`
	CSRFToken   string    `json:"csrf_token"`
}
