package events

type UserCreated struct {
	UserId   UserId
	Email    string
	Password string
}

func (u *UserCreated) ID() string {
	return "app.UserCreated"
}
