package port

import (
	"app/internal/core/component/user/application/queries/dto"
	"context"
)

type UserQueries interface {
	FindById(ctx context.Context, id string) (*dto.UserDTO, error)
}
