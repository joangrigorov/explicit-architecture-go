package ent

import (
	"app/internal/core/component/user/domain/user"
	"app/internal/core/component/user/domain/verification"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	roles "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"fmt"
)

func mapDomainRole(role roles.Role) user.Role {
	var domainRole user.Role
	switch role.String() {
	case "admin":
		domainRole = &user.Admin{}
	case "member":
		domainRole = &user.Member{}
	default:
		panic(fmt.Sprintf("unknown dto role %s", role.String()))
	}

	return domainRole
}

func mapDtoRole(role user.Role) roles.Role {
	var dtoRole roles.Role
	switch role.ID() {
	case "admin":
		dtoRole = roles.RoleAdmin
	case "member":
		dtoRole = roles.RoleMember
	default:
		panic(fmt.Sprintf("unknown domain role %s", dtoRole.String()))
	}

	return dtoRole
}

func mapUserAggregate(dto *generated.User) *user.User {
	if dto == nil {
		return nil
	}

	var idPUserId *user.IdPUserID
	if dto.IdpUserID != nil {
		tmp := user.IdPUserID(*dto.IdpUserID)
		idPUserId = &tmp
	} else {
		idPUserId = nil
	}

	return &user.User{
		ID:          user.ID(dto.ID.String()),
		Username:    user.Username(dto.Username),
		Email:       user.Email(dto.Email),
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		ConfirmedAt: dto.ConfirmedAt,
		Role:        mapDomainRole(dto.Role),
		IdPUserId:   idPUserId,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

func mapVerificationAggregate(dto *generated.Verification) (*verification.Verification, error) {
	if dto == nil {
		return nil, fmt.Errorf("ent: cannot map Verification dto - nil object passed")
	}

	csrfToken, err := verification.DecodeCSRFToken(dto.HashedToken)

	if err != nil {
		return nil, fmt.Errorf("decode csrf token: %w", err)
	}

	return &verification.Verification{
		ID:        verification.ID(dto.ID.String()),
		UserID:    user.ID(dto.UserID.String()),
		CSRFToken: csrfToken,
		ExpiresAt: dto.ExpiresAt,
		UsedAt:    dto.UsedAt,
		CreatedAt: dto.CreatedAt,
	}, nil
}
