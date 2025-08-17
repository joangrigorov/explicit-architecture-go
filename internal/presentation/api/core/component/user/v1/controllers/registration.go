package controllers

import (
	"app/internal/core/component/user/application/usecases"
	eventBus "app/internal/core/port/events"
	"app/internal/presentation/api/core/component/user/v1/requests"
	. "app/internal/presentation/api/core/component/user/v1/responses"
	. "app/internal/presentation/api/core/shared/responses"
	ctx "app/internal/presentation/api/port/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/jaswdr/faker/v2"
)

type RegistrationController struct {
	registrationUseCase *usecases.Registration
	confirmationUseCase *usecases.Confirmation
	eventBus            eventBus.EventBus
	tr                  ut.Translator
}

func NewRegistrationController(
	registration *usecases.Registration,
	confirmation *usecases.Confirmation,
	bus eventBus.EventBus,
	tr ut.Translator,
) *RegistrationController {
	return &RegistrationController{
		registrationUseCase: registration,
		confirmationUseCase: confirmation,
		eventBus:            bus,
		tr:                  tr,
	}
}

func (ctrl *RegistrationController) Register(ctx ctx.Context) {
	f := faker.New()

	user, err := ctrl.registrationUseCase.Execute(
		ctx.Context(),
		f.Internet().User(),
		f.Internet().Password(),
		f.Internet().Email(),
		f.Person().FirstName(),
		f.Person().LastName(),
	)

	if err != nil {
		InternalServerError(ctx, NewDefaultError(err))
		return
	}

	ctx.JSON(200, NewRegistrationResponse(user))
}

func (ctrl *RegistrationController) Confirm(ctx ctx.Context) {
	request := &requests.ConfirmationRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		UnprocessableEntity(ctx, Render(ctrl.tr, err, request))
		return
	}

	err := ctrl.confirmationUseCase.Execute(ctx.Context(), request.UserID)
	if err != nil {
		BadRequest(ctx, NewDefaultError(err))
		return
	}

	ctx.NoContent()
}
