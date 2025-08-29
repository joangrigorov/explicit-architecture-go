package find_user_by_id

import (
	"app/internal/core/port/cqrs"
	"context"
)

type Handler struct {
	queries UserQueries
}

func NewFindUserByIDHandler(queries UserQueries) cqrs.QueryHandler[Query, *UserDTO] {
	return &Handler{queries: queries}
}

func (h *Handler) Execute(ctx context.Context, q Query) (*UserDTO, error) {
	return h.queries.FindById(ctx, q.ID)
}
