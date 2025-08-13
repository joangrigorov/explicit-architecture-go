package domain

type RoleId string

type Role interface {
	Id() RoleId
	String() string
}

type Member struct{}

func (m *Member) String() string {
	return string(m.Id())
}

func (m *Member) Id() RoleId {
	return "member"
}

type Admin struct{}

func (a *Admin) String() string {
	return string(a.Id())
}

func (a *Admin) Id() RoleId {
	return "admin"
}
