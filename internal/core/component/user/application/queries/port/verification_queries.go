package port

import (
	"context"
)

type VerificationQueries interface {
	FindByID(ctx context.Context, id string) (*VerificationDTO, error)
	FindByUserID(ctx context.Context, userID string) (*VerificationDTO, error)
}
