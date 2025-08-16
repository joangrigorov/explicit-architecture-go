package domain

type RoleId string

func (roleId RoleId) String() string {
	return string(roleId)
}

type Role interface {
	ID() RoleId
}

type Member struct{}

func (m Member) ID() RoleId {
	return "member"
}

type Admin struct{}

func (a Admin) ID() RoleId {
	return "admin"
}
