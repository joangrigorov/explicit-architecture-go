package user

import (
	"app/internal/core/component/user/domain"
	"app/internal/core/shared_kernel/events"
	"app/internal/infrastructure/persistence/ent/generated/user"
	roles "app/internal/infrastructure/persistence/ent/generated/user/user"
	"fmt"
)

func mapDomainRole(role roles.Role) domain.Role {
	var domainRole domain.Role
	switch role.String() {
	case "admin":
		domainRole = &domain.Admin{}
	case "member":
		domainRole = &domain.Member{}
	default:
		panic(fmt.Sprintf("unknown dto role %s", role.String()))
	}

	return domainRole
}

func mapDtoRole(role domain.Role) roles.Role {
	var dtoRole roles.Role
	switch role.String() {
	case "admin":
		dtoRole = roles.RoleAdmin
	case "member":
		dtoRole = roles.RoleMember
	default:
		panic(fmt.Sprintf("unknown domain role %s", dtoRole.String()))
	}

	return dtoRole
}

func mapEntity(dto *user.User) *domain.User {
	idPUserId := domain.IdPUserId(*dto.IdpUserID)
	return domain.ReconstituteUser(
		events.UserId(dto.ID.String()),
		dto.Username,
		dto.Email,
		dto.FirstName,
		dto.LastName,
		mapDomainRole(dto.Role),
		&idPUserId,
		dto.ConfirmedAt,
		dto.CreatedAt,
		dto.UpdatedAt,
	)
}
