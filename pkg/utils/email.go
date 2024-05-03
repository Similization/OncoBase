package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

// EmailService provides functionality to send emails.
type EmailService struct {
	from     string // Sender email address
	password string // Sender email password
}

// NewEmailService creates a new EmailService instance.
func NewEmailService() *EmailService {
	return &EmailService{
		from:     "from@gmail.com",
		password: os.Getenv("EMAIL_PASSWORD"), // Get email password from environment variable
	}
}

// Authentication returns an smtp.Auth object for authentication.
func (es *EmailService) Authentication() smtp.Auth {
	return smtp.PlainAuth("", es.from, es.password, smtpHost)
}

// SendEmail sends an email with the provided subject and body to the specified recipients.
func (es *EmailService) SendEmail(to []string, subject, body string) error {
	// Construct email message with subject and body
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	// Construct SMTP server address
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	// Send email using SMTP server and authentication
	err := smtp.SendMail(addr, es.Authentication(), es.from, to, message)
	return err
}
