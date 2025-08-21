package ent

import (
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/infrastructure/component/user/persistence/ent/generated"
	roles "app/internal/infrastructure/component/user/persistence/ent/generated/user"
	"fmt"
)

func mapDomainRole(role roles.Role) Role {
	var domainRole Role
	switch role.String() {
	case "admin":
		domainRole = &Admin{}
	case "member":
		domainRole = &Member{}
	default:
		panic(fmt.Sprintf("unknown dto role %s", role.String()))
	}

	return domainRole
}

func mapDtoRole(role Role) roles.Role {
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

func mapEntity(dto *generated.User) *User {
	if dto == nil {
		return nil
	}

	var idPUserId *IdPUserID
	if dto.IdpUserID != nil {
		tmp := IdPUserID(*dto.IdpUserID)
		idPUserId = &tmp
	} else {
		idPUserId = nil
	}

	return ReconstituteUser(
		UserID(dto.ID.String()),
		dto.Username,
		dto.Email,
		dto.FirstName,
		dto.LastName,
		mapDomainRole(dto.Role),
		idPUserId,
		dto.ConfirmedAt,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}
