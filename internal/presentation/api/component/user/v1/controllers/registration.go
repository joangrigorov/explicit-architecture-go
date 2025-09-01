package controllers

import (
	"app/internal/core/component/user/application/queries/dto"
	qs "app/internal/core/component/user/application/queries/find_user_by_id"
	"app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/cqrs/queries"
	ctx "app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/user/v1/requests"
	"app/internal/presentation/api/shared/errors"

	"github.com/google/uuid"
)

type Registration struct {
	commandBus   cqrs.CommandBus
	queryBus     cqrs.QueryBus
	errorHandler *errors.Handler
}

func NewRegistration(
	commandBus cqrs.CommandBus,
	queryBus cqrs.QueryBus,
	errorHandler *errors.Handler,
) *Registration {
	return &Registration{commandBus: commandBus, queryBus: queryBus, errorHandler: errorHandler}
}

func (c *Registration) Register(ctx ctx.Context) {
	r := &requests.Registration{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		c.errorHandler.RenderRequestError(ctx, err, r)
		return
	}

	userID := uuid.New().String()
	cmd := r.NewRegisterUserCommand(userID)

	if err := c.commandBus.Dispatch(ctx.Context(), cmd); err != nil {
		c.errorHandler.RenderApplicationError(ctx, err)
		return
	}

	query := qs.Query{ID: userID}
	userDTO, err := queries.Execute[*dto.UserDTO](ctx.Context(), c.queryBus, query)

	if err != nil {
		c.errorHandler.RenderApplicationError(ctx, err)
		return
	}

	ctx.JSON(200, userDTO)
}
