package mail

type Mailer interface {
	Send(to, cc, bcc []string, from, subject, text string) error
}
