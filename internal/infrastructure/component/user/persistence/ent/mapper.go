package ent

import (
	"app/internal/core/component/user/domain/confirmation"
	"app/internal/core/component/user/domain/user"
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

	return user.ReconstituteUser(
		user.ID(dto.ID.String()),
		user.Username(dto.Username),
		user.Email(dto.Email),
		dto.FirstName,
		dto.LastName,
		mapDomainRole(dto.Role),
		idPUserId,
		dto.ConfirmedAt,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}

func mapConfirmationAggregate(dto *generated.Confirmation) *confirmation.Confirmation {
	if dto == nil {
		return nil
	}

	return confirmation.ReconstituteConfirmation(
		confirmation.ID(dto.ID.String()),
		user.ID(dto.UserID.String()),
		dto.HmacSecret,
		dto.CreatedAt,
	)
}
