package controllers

import (
	. "app/internal/core/component/user/application/commands"
	qs "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/dto"
	. "app/internal/core/port/cqrs"
	"app/internal/infrastructure/framework/cqrs/queries"
	ctx "app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/user/v1/requests"
	. "app/internal/presentation/api/shared/responses"

	ut "github.com/go-playground/universal-translator"
	"github.com/google/uuid"
)

type RegistrationController struct {
	commandBus CommandBus
	queryBus   QueryBus
	tr         ut.Translator
}

func NewRegistrationController(commandBus CommandBus, queryBus QueryBus, tr ut.Translator) *RegistrationController {
	return &RegistrationController{commandBus: commandBus, queryBus: queryBus, tr: tr}
}

func (c *RegistrationController) Register(ctx ctx.Context) {
	r := &requests.Registration{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		UnprocessableEntity(ctx, Render(c.tr, err, r))
		return
	}

	userID := uuid.New().String()

	cmd := NewRegisterUserCommand(
		userID,
		r.Username,
		r.Password,
		r.Email,
		r.FirstName,
		r.LastName,
	)

	if err := c.commandBus.Dispatch(ctx.Context(), cmd); err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	userDTO, err := queries.Execute[*dto.UserDTO](ctx.Context(), c.queryBus, qs.FindUserByIDQuery{ID: userID})

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, userDTO)
}

func (c *RegistrationController) Confirm(ctx ctx.Context) {
	request := &requests.Confirmation{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		UnprocessableEntity(ctx, Render(c.tr, err, request))
		return
	}

	err := c.commandBus.Dispatch(ctx.Context(), NewConfirmUserCommand(request.UserID))

	if err != nil {
		BadRequest(ctx, NewDefaultError(err))
		return
	}

	userDTO, err := queries.Execute[*dto.UserDTO](ctx.Context(), c.queryBus, qs.FindUserByIDQuery{ID: request.UserID})

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, userDTO)
}
