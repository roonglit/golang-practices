package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"

	"github.com/hibiken/asynq"
)

const TypeEmailSend = "email:send"

type EmailPayload struct {
	To      string
	Subject string
	Body    string
}

func NewEmailTask(to, subject, body string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailPayload{To: to, Subject: subject, Body: body})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailSend, payload), nil
}

func HandleEmailTask(ctx context.Context, t *asynq.Task) error {
	var p EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	err := sendEmail(p.To, p.Subject, p.Body)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Email sent to %s", p.To)
	return nil
}

func sendEmail(to string, subject string, body string) error {
	// Sender data.
	from := "sitthidech.p@gmail.com"
	password := "qhzp nbyu bgub udjo" // Use App Password if using Gmail -> https://myaccount.google.com/apppasswords

	// Receiver email address.
	tos := []string{to}

	// SMTP server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message body.
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, tos, []byte(message))
	if err != nil {
		return err
	}
	fmt.Println("Email Sent Successfully!")

	return nil
}
