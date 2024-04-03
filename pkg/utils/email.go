package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

const (
	from     = "from@gmail.com"
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"

	messageRegister       = "Your code to register:"
	messagePasswordUpdate = "Your code to update your password:"
	messageIgnore         = "If it was not you who made the request - ignore this message."
)

type EmailService struct {
	from     string
	host     string
	port     string
	password string
}

func (es *EmailService) NewEmailService() *EmailService {
	return &EmailService{
		from:     from,
		host:     smtpHost,
		port:     smtpPort,
		password: os.Getenv("EMAIL_PASSWORD"),
	}
}

func (es *EmailService) Authentication() smtp.Auth {
	return smtp.PlainAuth("", es.from, es.password, es.host)
}

func createMessage(txt string) string {
	return txt
}

func (es *EmailService) SendEmail(auth smtp.Auth, to []string, message []byte) error {
	addr := fmt.Sprintf("%s:%s", es.host, es.port)
	err := smtp.SendMail(addr, auth, es.from, to, message)
	return err
}
