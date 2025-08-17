package controllers

import (
	. "app/internal/core/component/user/application/commands"
	"app/internal/core/port/commands"
	"app/internal/presentation/api/core/component/user/v1/requests"
	. "app/internal/presentation/api/core/shared/responses"
	ctx "app/internal/presentation/api/port/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/jaswdr/faker/v2"
)

type RegistrationController struct {
	commandBus commands.CommandBus
	tr         ut.Translator
}

func NewRegistrationController(commandBus commands.CommandBus, tr ut.Translator) *RegistrationController {
	return &RegistrationController{commandBus: commandBus, tr: tr}
}

func (c *RegistrationController) Register(ctx ctx.Context) {
	f := faker.New()

	cmd := NewRegisterUserCommand(
		f.Internet().User(),
		f.Internet().Password(),
		f.Internet().Email(),
		f.Person().FirstName(),
		f.Person().LastName(),
	)

	err := c.commandBus.Dispatch(ctx.Context(), cmd)

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.NoContent()
}

func (c *RegistrationController) Confirm(ctx ctx.Context) {
	request := &requests.ConfirmationRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		UnprocessableEntity(ctx, Render(c.tr, err, request))
		return
	}

	err := c.commandBus.Dispatch(ctx.Context(), NewConfirmUserCommand(request.UserID))
	
	if err != nil {
		BadRequest(ctx, NewDefaultError(err))
		return
	}

	ctx.NoContent()
}
