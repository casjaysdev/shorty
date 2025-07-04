// File: internal/lib/utils/email.go
// Purpose: Handles email sending with support for templates, fallback SMTP, and scoped headers.

package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"text/template"
)

type EmailConfig struct {
	FromName     string
	FromAddress  string
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
}

func SendEmail(cfg EmailConfig, to []string, subject string, tpl string, data any) error {
	body, err := renderTemplate(tpl, data)
	if err != nil {
		return fmt.Errorf("render email: %w", err)
	}

	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPassword, cfg.SMTPHost)
	msg := buildMessage(cfg.FromName, cfg.FromAddress, to, subject, body)
	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)

	return smtp.SendMail(addr, auth, cfg.FromAddress, to, msg)
}

func renderTemplate(tpl string, data any) (string, error) {
	t, err := template.New("email").Parse(tpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func buildMessage(fromName, fromAddr string, to []string, subject, body string) []byte {
	var msg strings.Builder
	msg.WriteString(fmt.Sprintf("From: %s <%s>\r\n", fromName, fromAddr))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ",")))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(body)
	return []byte(msg.String())
}
