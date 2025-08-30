package find_user_by_id

import (
	"context"
)

type UserQueries interface {
	FindByID(ctx context.Context, id string) (*DTO, error)
}
