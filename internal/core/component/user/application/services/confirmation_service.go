package services

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain"
	"app/internal/core/port/hmac"
	"app/internal/core/port/uuid"
	sk "app/internal/core/shared_kernel/domain"
	"context"
)

type ConfirmationService struct {
	confirmRepository repositories.ConfirmationRepository
	uuidGenerator     uuid.Generator
	hmacGen           hmac.Generator
}

func NewConfirmationService(
	confirmRepository repositories.ConfirmationRepository,
	uuidGenerator uuid.Generator,
	hmacGen hmac.Generator,
) *ConfirmationService {
	return &ConfirmationService{
		confirmRepository: confirmRepository,
		uuidGenerator:     uuidGenerator,
		hmacGen:           hmacGen,
	}
}

func (s *ConfirmationService) Create(ctx context.Context, userID sk.UserID) (
	confirmation *domain.Confirmation,
	hmac *string,
	err error,
) {
	id := domain.ConfirmationID(s.uuidGenerator.Generate())
	hmacSum, secret, err := s.hmacGen.Generate(id.String())

	if err != nil {
		return nil, nil, err
	}

	c := domain.NewConfirmation(id, userID, secret)

	if err := s.confirmRepository.Create(ctx, c); err != nil {
		return nil, nil, err
	}

	return c, &hmacSum, nil
}
