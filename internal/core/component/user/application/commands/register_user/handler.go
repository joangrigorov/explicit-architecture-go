package register_user

import (
	"app/internal/core/component/user/application/repositories"
	. "app/internal/core/component/user/domain/user"
	"app/internal/core/port/errors"
	"context"
)

type Handler struct {
	userRepository repositories.UserRepository
	errors         errors.ErrorFactory
}

func NewHandler(
	userRepository repositories.UserRepository,
	errors errors.ErrorFactory,
) *Handler {
	return &Handler{
		userRepository: userRepository,
		errors:         errors,
	}
}

func (h *Handler) Handle(ctx context.Context, c Command) error {
	id := ID(c.userID)

	if user, _ := h.userRepository.GetById(ctx, id); user != nil {
		return h.errors.New(errors.ErrConflict, "User with that ID already exists", nil)
	}

	email, err := NewEmail(c.email)
	if err != nil {
		return h.errors.New(errors.ErrValidation, "Invalid email format", err)
	}

	if user, _ := h.userRepository.GetByEmail(ctx, email); user != nil {
		return h.errors.New(errors.ErrConflict, "User with that email already exists", nil)
	}

	username, err := NewUsername(c.username)
	if err != nil {
		return h.errors.New(errors.ErrValidation, "Invalid username provided", err)
	}

	if user, _ := h.userRepository.GetByUsername(ctx, username); user != nil {
		return h.errors.New(errors.ErrConflict, "User with that username already exists", nil)
	}

	user := NewUser(id, username, email, c.firstName, c.lastName, &Member{})

	if err := h.userRepository.Create(ctx, user); err != nil {
		return h.errors.New(errors.ErrDB, "Error creating user", err)
	}

	return nil
}
