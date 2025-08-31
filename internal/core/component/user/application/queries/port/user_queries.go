package port

import (
	"app/internal/core/component/user/application/queries/dto"
	"context"
)

type UserQueries interface {
	FindByID(ctx context.Context, id string) (*dto.UserDTO, error)
}
