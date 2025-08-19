package queries

import (
	. "app/internal/core/component/user/application/queries/port"
	"app/internal/core/port/cqrs"
	"context"
	"encoding/json"
)

type FindUserByIDQuery struct {
	ID string
}

func (q FindUserByIDQuery) LogBody() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id": q.ID,
	})
}

type FindUserByIDHandler struct {
	queries UserQueries
}

func NewFindUserByIDHandler(queries UserQueries) cqrs.QueryHandler[FindUserByIDQuery, *UserDTO] {
	return &FindUserByIDHandler{queries: queries}
}

func (h *FindUserByIDHandler) Execute(ctx context.Context, q FindUserByIDQuery) (*UserDTO, error) {
	return h.queries.FindById(ctx, q.ID)
}
