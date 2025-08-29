package user

import "regexp"

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Email string

func NewEmail(email string) (Email, error) {
	e := Email(email)
	if !e.IsValid() {
		return "", InvalidEmailError{}
	}

	return e, nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) IsValid() bool {
	return emailRegex.MatchString(string(e))
}

type InvalidEmailError struct{}

func (i InvalidEmailError) Error() string {
	return "Email format is invalid"
}
