package mailables

type UserConfirmedMail interface {
	Render(fullName string) (message string, err error)
}
