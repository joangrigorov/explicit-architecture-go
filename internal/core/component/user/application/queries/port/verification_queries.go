package port

import (
	"app/internal/core/component/user/application/queries/dto"
	"context"
)

type VerificationQueries interface {
	FindByID(ctx context.Context, id string) (*dto.VerificationDTO, error)
	FindByUserID(ctx context.Context, userID string) (*dto.VerificationDTO, error)
}
