package port

import "context"

type UserQueries interface {
	FindById(ctx context.Context, id string) (*UserDTO, error)
}
