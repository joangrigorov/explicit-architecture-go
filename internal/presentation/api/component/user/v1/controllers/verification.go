package controllers

import (
	"app/internal/core/component/user/application/commands/complete_password_setup"
	"app/internal/core/component/user/application/queries/dto"
	preflight_check "app/internal/core/component/user/application/queries/get_verification_preflight"
	"app/internal/core/port/cqrs"
	errorsPort "app/internal/core/port/errors"
	"app/internal/core/port/logging"
	"app/internal/infrastructure/framework/cqrs/queries"
	"app/internal/infrastructure/framework/http"
	"app/internal/presentation/api/component/user/v1/requests"
	"app/internal/presentation/api/shared/responses"
	"errors"

	ut "github.com/go-playground/universal-translator"
)

type Verification struct {
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
	logger     logging.Logger
	tr         ut.Translator
}

func NewVerification(
	commandBus cqrs.CommandBus,
	queryBus cqrs.QueryBus,
	logger logging.Logger,
	tr ut.Translator,
) *Verification {
	return &Verification{
		commandBus: commandBus,
		queryBus:   queryBus,
		logger:     logger,
		tr:         tr,
	}
}

func (v *Verification) PreflightValidate(ctx http.Context) {
	verificationID := ctx.ParamString("id")
	token := ctx.Query("token")

	if verificationID == "" {
		responses.UnprocessableEntity(ctx, responses.DefaultError{Error: "Verification id is required"})
		return
	}

	if token == "" {
		responses.UnprocessableEntity(ctx, responses.DefaultError{Error: "Token is required"})
		return
	}

	query := preflight_check.NewQuery(verificationID, token)
	result, err := queries.Execute[*dto.PreflightDTO](ctx.Context(), v.queryBus, query)
	if err != nil {
		v.renderVerificationError(ctx, err)
		return
	}

	if !result.ValidCSRF {
		responses.BadRequest(ctx, result)
		return
	}

	if result.Expired {
		responses.Gone(ctx, result)
		return
	}

	ctx.JSON(200, result)
}

func (v *Verification) PasswordSetup(ctx http.Context) {
	req := &requests.PasswordSetup{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.UnprocessableEntity(ctx, responses.Render(v.tr, err, req))
		return
	}
	verificationID := ctx.ParamString("id")
	token := req.Token

	query := preflight_check.NewQuery(verificationID, token)
	result, err := queries.Execute[*dto.PreflightDTO](ctx.Context(), v.queryBus, query)

	if err != nil {
		v.renderVerificationError(ctx, err)
		return
	}

	if !result.ValidCSRF {
		responses.BadRequest(ctx, result)
		return
	}

	if result.Expired {
		responses.Gone(ctx, result)
		return
	}

	cmd := complete_password_setup.NewCommand(result.UserID, req.Password)

	if err = v.commandBus.Dispatch(ctx.Context(), cmd); err != nil {
		// TODO handle other types of errors, not all is 500
		responses.InternalServerError(ctx, responses.NewDefaultError(err))
		return
	}

	ctx.NoContent()
}

func (v *Verification) renderVerificationError(ctx http.Context, err error) {
	if err != nil {
		var appErr errorsPort.Error
		if errors.As(err, &appErr) {
			switch appErr.Code() {
			case errorsPort.ErrNotFound:
				responses.NotFound(ctx, responses.NewDefaultError(err))
			case errorsPort.ErrValidation:
				responses.UnprocessableEntity(ctx, responses.NewDefaultError(err))
			case errorsPort.ErrConflict:
				responses.Conflict(ctx, responses.NewDefaultError(err))
			case errorsPort.ErrDB:
				responses.InternalServerError(ctx, responses.NewDefaultError(err))
			default:
				responses.InternalServerError(ctx, responses.DefaultError{Error: "Unknown error occurred"})
			}
		} else {
			v.logger.Error(err)
			responses.InternalServerError(ctx, responses.DefaultError{Error: "Unknown error occurred"})
		}
	}
}
