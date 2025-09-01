package controllers

import (
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/find_user_by_id"
	"app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/shared/errors"
)

type Me struct {
	queryBus     cqrs.QueryBus
	errorHandler *errors.Handler
}

func (m *Me) Me(c http.Context) {
	query := find_user_by_id.NewQuery(c.GetHeader("x-app-user-id"))
	res, err := m.queryBus.Execute(c.Context(), query)
	if err != nil {
		m.errorHandler.RenderApplicationError(c, err)
		return
	}

	c.JSON(200, res.(*dto.UserDTO))
}

func NewMe(queryBus cqrs.QueryBus, errorHandler *errors.Handler) *Me {
	return &Me{queryBus: queryBus, errorHandler: errorHandler}
}
