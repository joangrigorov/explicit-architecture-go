package get_verification_preflight

import (
	"app/internal/core/component/user/domain/verification"
	"app/internal/core/port/errors"
	"context"
	"crypto/subtle"
)

type Handler struct {
	queries VerificationQueries
	errors  errors.ErrorFactory
}

func NewHandler(queries VerificationQueries, errors errors.ErrorFactory) *Handler {
	return &Handler{queries: queries, errors: errors}
}

func (h *Handler) Execute(ctx context.Context, q Query) (*VerificationPreflightDto, error) {
	ver, err := h.queries.FindByID(ctx, q.verificationID)
	if err != nil || ver == nil {
		return nil, h.errors.New(errors.ErrDB, "Verification not found", err)
	}

	expectedCSRF := verification.Token(q.Token).Hash()
	actualCSRF, err := verification.DecodeCSRFToken(ver.CSRFToken)

	if err != nil {
		return nil, h.errors.New(errors.ErrDB, "Cannot decode stored CSRF token", err)
	}

	// TODO continue from here
	if subtle.ConstantTimeCompare([]byte(expectedCSRF.Bytes()), actualCSRF.Bytes()) == 1
}
