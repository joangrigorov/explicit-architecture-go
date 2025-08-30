package services

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/component/user/domain/verification"
	"app/internal/core/port/errors"
	"app/internal/core/port/uuid"
	"context"
)

type VerificationService struct {
	verificationRepository repositories.VerificationRepository
	uuidGenerator          uuid.Generator
	errors                 errors.ErrorFactory
}

func NewVerificationService(
	verificationRepository repositories.VerificationRepository,
	uuidGenerator uuid.Generator,
	errors errors.ErrorFactory,
) *VerificationService {
	return &VerificationService{
		verificationRepository: verificationRepository,
		uuidGenerator:          uuidGenerator,
		errors:                 errors,
	}
}

func (s *VerificationService) Create(ctx context.Context, userID user.ID, userEmail user.Email) (
	ver *verification.Verification,
	token string,
	err error,
) {
	id := verification.ID(s.uuidGenerator.Generate())
	t, err := verification.GenerateToken()

	if err != nil {
		return nil, "", s.errors.New(errors.ErrUnknown, "Error generating verification token", err)
	}

	c := verification.NewVerification(id, userID, userEmail.Mask(), t.Hash())

	if err := s.verificationRepository.Create(ctx, c); err != nil {
		return nil, "", s.errors.New(errors.ErrUnknown, "Error creating verification", err)
	}

	return c, t.Encode(), nil
}
