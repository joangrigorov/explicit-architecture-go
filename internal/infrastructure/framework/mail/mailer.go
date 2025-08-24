package mail

import (
	port "app/internal/core/port/mail"

	"gopkg.in/mail.v2"
)

type Mailer struct {
	dialer *mail.Dialer
}

func NewGomailMailer(dialer *mail.Dialer) *Mailer {
	return &Mailer{dialer: dialer}
}

func NewMailer(m *Mailer) port.Mailer {
	return m
}

func (m *Mailer) Send(to, cc, bcc []string, from, subject, text string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Cc", cc...)
	msg.SetHeader("Bcc", bcc...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", text)

	return m.dialer.DialAndSend(msg)
}
