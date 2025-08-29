package controllers

import (
	qs "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/cqrs/queries"
	ctx "app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/user/v1/requests"
	. "app/internal/presentation/api/shared/responses"

	ut "github.com/go-playground/universal-translator"
	"github.com/google/uuid"
)

type RegistrationController struct {
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
	tr         ut.Translator
}

func NewRegistrationController(commandBus cqrs.CommandBus, queryBus cqrs.QueryBus, tr ut.Translator) *RegistrationController {
	return &RegistrationController{commandBus: commandBus, queryBus: queryBus, tr: tr}
}

func (c *RegistrationController) Register(ctx ctx.Context) {
	r := &requests.Registration{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		UnprocessableEntity(ctx, Render(c.tr, err, r))
		return
	}

	userID := uuid.New().String()
	cmd := r.NewRegisterUserCommand(userID)

	if err := c.commandBus.Dispatch(ctx.Context(), cmd); err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	query := qs.FindUserByIDQuery{ID: userID}
	userDTO, err := queries.Execute[*dto.UserDTO](ctx.Context(), c.queryBus, query)

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, userDTO)
}
