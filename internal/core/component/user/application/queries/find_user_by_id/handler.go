package find_user_by_id

import (
	"app/internal/core/port/cqrs"
	"context"
)

type Handler struct {
	queries UserQueries
}

func NewHandler(queries UserQueries) cqrs.QueryHandler[Query, *DTO] {
	return &Handler{queries: queries}
}

func (h *Handler) Execute(ctx context.Context, q Query) (*DTO, error) {
	return h.queries.FindByID(ctx, q.ID)
}
