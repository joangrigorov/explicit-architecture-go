package controllers

import (
	. "app/internal/core/component/user/application/commands"
	qs "app/internal/core/component/user/application/queries"
	"app/internal/core/component/user/application/queries/port"
	. "app/internal/core/port/commands"
	. "app/internal/core/port/queries"
	"app/internal/infrastructure/queries"
	"app/internal/presentation/api/core/component/user/v1/requests"
	. "app/internal/presentation/api/core/shared/responses"
	ctx "app/internal/presentation/api/port/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
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
	f := faker.New()

	userID := uuid.New().String()

	cmd := NewRegisterUserCommand(
		userID,
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

	userDTO, err := queries.Execute[*port.UserDTO](ctx.Context(), c.queryBus, qs.FindUserByIDQuery{ID: userID})

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, userDTO)
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

	userDTO, err := queries.Execute[*port.UserDTO](ctx.Context(), c.queryBus, qs.FindUserByIDQuery{ID: request.UserID})

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, userDTO)
}
