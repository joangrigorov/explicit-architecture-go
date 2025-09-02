package register_user

import (
	"app/internal/core/component/user/application/repositories"
	"app/internal/core/component/user/domain/user"
	"app/internal/core/port/errors"
	"context"
)

type Handler struct {
	userRepository repositories.UserRepository
	errors         errors.ErrorFactory
}

func (h *Handler) Handle(ctx context.Context, c Command) error {
	id := user.ID(c.userID)

	if usr, _ := h.userRepository.GetById(ctx, id); usr != nil {
		return h.errors.New(errors.ErrConflict, "User with that ID already exists", nil)
	}

	email, err := user.NewEmail(c.email)
	if err != nil {
		return h.errors.New(errors.ErrValidation, "Invalid email format", err)
	}

	if usr, _ := h.userRepository.GetByEmail(ctx, email); usr != nil {
		return h.errors.New(errors.ErrConflict, "User with that email already exists", nil)
	}

	username, err := user.NewUsername(c.username)
	if err != nil {
		return h.errors.New(errors.ErrValidation, "Invalid username provided", err)
	}

	if usr, _ := h.userRepository.GetByUsername(ctx, username); usr != nil {
		return h.errors.New(errors.ErrConflict, "User with that username already exists", nil)
	}

	usr := user.NewUser(id, username, email, c.firstName, c.lastName, user.Member{})

	if err := h.userRepository.Create(ctx, usr); err != nil {
		return h.errors.New(errors.ErrDB, "Error creating user", err)
	}

	return nil
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
