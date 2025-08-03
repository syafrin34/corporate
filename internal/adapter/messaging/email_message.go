package messaging

import (
	"corporate/config"
	"crypto/tls"

	"github.com/go-mail/mail"
	"github.com/labstack/gommon/log"
)

type EmailMessagingInterface interface {
	SendEMailAppointment(attach *string, from, subject, body string) error
}
type emailAttribut struct {
	username string
	password string
	host     string
	port     int
	isTLS    bool
	receiver string
}

// SendEMailAppointment implements EmailMessagingInterface.
func (e *emailAttribut) SendEMailAppointment(attach *string, from, subject string, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", e.receiver)

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if attach != nil {
		m.Attach(*attach)
	}
	d := mail.NewDialer(e.host, e.port, e.username, e.password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: e.isTLS,
	}
	if err := d.DialAndSend(m); err != nil {
		log.Errorf("error sending mail: %v", err)
		return err
	}
	return nil
}

func NewEmailMessaging(cfg *config.Config) EmailMessagingInterface {
	return &emailAttribut{
		username: cfg.Email.Username,
		password: cfg.Email.Password,
		host:     cfg.Email.Host,
		port:     cfg.Email.Port,
		isTLS:    cfg.Email.IsTLS,
		receiver: cfg.Email.Receiver,
	}

}
