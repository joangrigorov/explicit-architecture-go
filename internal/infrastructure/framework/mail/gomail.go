package mail

import (
	"app/config/api"

	"gopkg.in/mail.v2"
)

func NewDialer(cfg api.Config) *mail.Dialer {
	return mail.NewDialer(cfg.Mail.Host, cfg.Mail.Port, cfg.Mail.User, cfg.Mail.Password)
}
