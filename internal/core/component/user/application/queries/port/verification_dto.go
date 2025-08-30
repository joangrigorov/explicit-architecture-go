package port

import "time"

type VerificationDTO struct {
	ID              string
	UserID          string
	UserEmailMasked string
	CSRFToken       string
	ExpiresAt       time.Time
	UsedAt          *time.Time
	CreatedAt       time.Time
}
