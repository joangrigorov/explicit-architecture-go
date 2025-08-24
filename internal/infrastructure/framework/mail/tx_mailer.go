package mail

import (
	"app/internal/core/port/logging"
)

type mailable struct {
	to, cc, bcc         []string
	from, subject, text string
}

type TransactionalMailer struct {
	mailer *Mailer
	logger logging.Logger
	outbox []mailable
}

func NewTransactionalMailer(m *Mailer, l logging.Logger) *TransactionalMailer {
	return &TransactionalMailer{mailer: m, logger: l}
}

func (m *TransactionalMailer) Send(to, cc, bcc []string, from, subject, text string) error {
	m.outbox = append(m.outbox, mailable{to: to, cc: cc, bcc: bcc, from: from, subject: subject, text: text})
	return nil
}

func (m *TransactionalMailer) Flush() {
	for _, ml := range m.outbox {
		err := m.mailer.Send(ml.to, ml.cc, ml.bcc, ml.from, ml.subject, ml.text)
		if err != nil {
			m.logger.Error("Could not send email: " + err.Error())
		}
	}
	m.Reset()
}

func (m *TransactionalMailer) Reset() {
	m.outbox = m.outbox[:0]
}
