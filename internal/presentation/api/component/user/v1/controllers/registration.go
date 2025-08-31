package controllers

import (
	"app/internal/core/component/user/application/queries/dto"
	qs "app/internal/core/component/user/application/queries/find_user_by_id"
	"app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/cqrs/queries"
	ctx "app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/user/v1/requests"
	. "app/internal/presentation/api/shared/responses"

	ut "github.com/go-playground/universal-translator"
	"github.com/google/uuid"
)

type Registration struct {
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
	tr         ut.Translator
}

func NewRegistrationController(commandBus cqrs.CommandBus, queryBus cqrs.QueryBus, tr ut.Translator) *Registration {
	return &Registration{commandBus: commandBus, queryBus: queryBus, tr: tr}
}

func (c *Registration) Register(ctx ctx.Context) {
	r := &requests.Registration{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		UnprocessableEntity(ctx, Render(c.tr, err, r))
		return
	}

	userID := uuid.New().String()
	cmd := r.NewRegisterUserCommand(userID)

	if err := c.commandBus.Dispatch(ctx.Context(), cmd); err != nil {
		// TODO handle other types of errors. Not all is 500
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	query := qs.Query{ID: userID}
	userDTO, err := queries.Execute[*dto.UserDTO](ctx.Context(), c.queryBus, query)

	if err != nil {
		// TODO handle other types of errors. Not all is 500
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, userDTO)
}
