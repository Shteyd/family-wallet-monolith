package adapter

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"monolith/internal/module/mailing/core"
	"os/exec"
	"strings"
)

type SmtpAdapter interface {
	Ping() bool
	SendMail(mail core.Email) error
}

type _SmtpAdapter struct {
	client *gomail.Dialer
}

func NewAdapter() SmtpAdapter {
	dialer := gomail.NewDialer(core.SmtpHost, core.SmtpPort, core.Username, core.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &_SmtpAdapter{
		client: dialer,
	}
}

func (s _SmtpAdapter) Ping() bool {
	out, _ := exec.Command("ping", s.client.Host, "-c 5", "-i 3", "-w 10").Output()
	return strings.Contains(string(out), "Destination Host Unreachable")
}

func (s _SmtpAdapter) SendMail(mail core.Email) error {
	if !s.Ping() {
		return errors.New("smtp server are unavailable")
	}

	message := gomail.NewMessage()
	message.SetHeader(core.MessageHeaderFrom, core.Username)
	message.SetHeader(core.MessageHeaderTo, mail.ClientEmail)
	message.SetAddressHeader(core.MessageAddressHeaderCc, mail.ClientEmail, mail.ClientName)
	message.SetHeader(core.MessageHeaderSubject, mail.MessageHeader)
	message.SetBody(core.MessageBodyType, mail.MessageBody)

	return s.client.DialAndSend(message)

}
