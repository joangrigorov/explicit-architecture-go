package get_verification_preflight

import "time"

type VerificationDTO struct {
	ID        string
	UserID    string
	CSRFToken string
	ExpiresAt time.Time
	UsedAt    *time.Time
	CreatedAt time.Time
}
