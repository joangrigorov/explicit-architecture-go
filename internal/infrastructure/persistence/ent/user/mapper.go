package user

import (
	. "app/internal/core/component/user/domain"
	. "app/internal/core/shared_kernel/domain"
	"app/internal/infrastructure/persistence/ent/generated/user"
	roles "app/internal/infrastructure/persistence/ent/generated/user/user"
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

func mapEntity(dto *user.User) *User {
	var idPUserId *IdPUserId
	if dto != nil && dto.IdpUserID != nil {
		tmp := IdPUserId(*dto.IdpUserID)
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
