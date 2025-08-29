package services

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/confirmation"
	sk "app/internal/core/component/user/domain/user"
	"app/internal/core/port/hmac"
	"app/internal/core/port/uuid"
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

func (s *ConfirmationService) Create(ctx context.Context, userID sk.ID) (
	con *confirmation.Confirmation,
	hmac *string,
	err error,
) {
	id := confirmation.ID(s.uuidGenerator.Generate())
	hmacSum, secret, err := s.hmacGen.Generate(id.String())

	if err != nil {
		return nil, nil, err
	}

	c := confirmation.NewConfirmation(id, userID, secret)

	if err := s.confirmRepository.Create(ctx, c); err != nil {
		return nil, nil, err
	}

	return c, &hmacSum, nil
}
