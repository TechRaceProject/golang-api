// services/send_email.go
package services

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"api/src/config"

	"gopkg.in/mail.v2"
)

type EmailRequest struct {
	To       string
	Subject  string
	Template string
	Data     map[string]interface{}
}

func SendEmail(request EmailRequest, cfg *config.Config) error {
	templatePath := os.Getenv("TEMPLATE_PATH")
	fullPath := filepath.Join(templatePath, request.Template)

	log.Printf("Loading email template from: %s", fullPath)

	t, err := template.ParseFiles(fullPath)
	if err != nil {
		log.Printf("Error loading template: %v", err)
		return err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, request.Data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		return err
	}

	log.Printf("Sending email to: %s with subject: %s", request.To, request.Subject)

	m := mail.NewMessage()
	m.SetHeader("From", cfg.EmailFrom)
	m.SetHeader("To", request.To)
	m.SetHeader("Subject", request.Subject)
	m.SetBody("text/html", buf.String())

	d := mail.NewDialer(cfg.SmtpHost, 587, cfg.SmtpUser, cfg.SmtpPass)

	// Log SMTP configuration avant l'envoi
	log.Printf("SMTP Host: %s, SMTP User: %s", cfg.SmtpHost, cfg.SmtpUser)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Println("Email sent successfully")

	return nil
}
