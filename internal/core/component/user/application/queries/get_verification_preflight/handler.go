package get_verification_preflight

import (
	"app/internal/core/component/user/application/queries/dto"
	"app/internal/core/component/user/application/queries/port"
	"app/internal/core/component/user/domain/verification"
	"app/internal/core/port/errors"
	"context"
	"crypto/subtle"
	"time"
)

type Handler struct {
	verificationQueries port.VerificationQueries
	errors              errors.ErrorFactory
}

func NewHandler(queries port.VerificationQueries, errors errors.ErrorFactory) *Handler {
	return &Handler{verificationQueries: queries, errors: errors}
}

func (h *Handler) Execute(ctx context.Context, q Query) (*dto.PreflightDTO, error) {
	ver, err := h.verificationQueries.FindByID(ctx, q.verificationID)

	if err != nil || ver == nil {
		return nil, h.errors.New(errors.ErrNotFound, "Verification entry not found", err)
	}

	if ver.UsedAt != nil {
		return nil, h.errors.New(errors.ErrConflict, "Verification entry already used", nil)
	}

	rawToken, err := verification.DecodeToken(q.Token)
	if err != nil {
		return nil, h.errors.New(errors.ErrValidation, "Verification token decode failed", err)
	}

	expectedCSRF := rawToken.Hash()
	actualCSRF, err := verification.DecodeCSRFToken(ver.CSRFToken)

	if err != nil {
		return nil, h.errors.New(errors.ErrDB, "Cannot decode stored CSRF token", err)
	}

	return &dto.PreflightDTO{
		ValidCSRF:   subtle.ConstantTimeCompare(expectedCSRF[:], actualCSRF[:]) == 1,
		Expired:     time.Now().After(ver.ExpiresAt),
		MaskedEmail: ver.UserEmailMasked,
		UserID:      ver.UserID,
	}, nil
}
