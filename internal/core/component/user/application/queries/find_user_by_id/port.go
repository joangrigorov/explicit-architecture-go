package find_user_by_id

import (
	"context"
)

type UserQueries interface {
	FindById(ctx context.Context, id string) (*UserDTO, error)
}
