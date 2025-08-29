package get_verification_preflight

import (
	"context"
)

type VerificationQueries interface {
	FindByID(ctx context.Context, id string) (*VerificationDTO, error)
}
