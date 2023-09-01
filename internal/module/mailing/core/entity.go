package core

import "fmt"

type Email struct {
	ClientEmail   string
	ClientName    string
	MessageHeader string
	MessageBody   string
}

type ConfirmationMail struct {
	Email
	ConfirmationCode int
}

func NewEmail(client, clientName, messageBody, messageHeader string) Email {
	return Email{
		ClientEmail: client,
		ClientName:  clientName,
		MessageBody: messageBody,
	}
}

func NewConfirmationCode(client, clientName, messageBody, messageHeader string, confirmationCode int) ConfirmationMail {
	messageBody = fmt.Sprintf(ConfirmCodeBodyTemplate, clientName, confirmationCode)
	return ConfirmationMail{
		Email:            NewEmail(client, clientName, messageBody, messageHeader),
		ConfirmationCode: confirmationCode,
	}
}

func (s *ConfirmationMail) GetDefaultEmail() Email {
	return s.Email
}

func (s *ConfirmationMail) GetConfirmationCode() int {
	return s.ConfirmationCode
}
