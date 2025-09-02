package find_user_by_id

import (
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/cqrs"
	"context"
)

type Handler struct {
	queries port.UserQueries
}

func (h *Handler) Execute(ctx context.Context, q Query) (*dto.UserDTO, error) {
	return h.queries.FindByID(ctx, q.ID)
}

func NewHandler(queries port.UserQueries) cqrs.QueryHandler[Query, *dto.UserDTO] {
	return &Handler{queries: queries}
}
