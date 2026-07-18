package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"mime"
	"net"
	"net/smtp"
	"strings"
	"time"

	"backend.app/configs"
	"backend.app/internal/models"
)

type SMTP struct{ config *configs.Config }

func NewSMTP(config *configs.Config) EmailService { return &SMTP{config: config} }

func (m *SMTP) SendEmail(job models.NotificationJob) error {
	body, err := parseEmailTemplate(job)
	if err != nil {
		return err
	}
	subject, _ := job.Content["subject"].(string)
	if subject == "" {
		subject = "Your Celebut verification code"
	}
	from := job.From
	if from == "" {
		from = m.config.MailFromAddress
	}
	fromName := job.FromName
	if fromName == "" {
		fromName = m.config.MailFromName
	}
	var message bytes.Buffer
	fmt.Fprintf(&message, "From: %s <%s>\r\n", mime.QEncoding.Encode("UTF-8", fromName), from)
	fmt.Fprintf(&message, "To: %s\r\n", strings.Join(job.To, ", "))
	fmt.Fprintf(&message, "Subject: %s\r\n", mime.QEncoding.Encode("UTF-8", subject))
	message.WriteString("MIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n")
	message.WriteString(*body)
	auth := smtp.PlainAuth("", m.config.MailUsername, m.config.MailPassword, m.config.MailHost)
	address := net.JoinHostPort(m.config.MailHost, m.config.MailPort)
	if m.config.MailEncryption == "ssl" || m.config.MailPort == "465" {
		return sendSMTPOverTLS(address, m.config.MailHost, auth, from, job.To, message.Bytes())
	}
	// smtp.SendMail automatically upgrades with STARTTLS when the server advertises it.
	return smtp.SendMail(address, auth, from, job.To, message.Bytes())
}

func sendSMTPOverTLS(address, host string, auth smtp.Auth, from string, recipients []string, message []byte) error {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 12 * time.Second}, "tcp", address, &tls.Config{ServerName: host, MinVersion: tls.VersionTLS12})
	if err != nil {
		return err
	}
	defer conn.Close()
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer client.Close()
	if err = client.Auth(auth); err != nil {
		return err
	}
	if err = client.Mail(from); err != nil {
		return err
	}
	for _, recipient := range recipients {
		if err = client.Rcpt(recipient); err != nil {
			return err
		}
	}
	writer, err := client.Data()
	if err != nil {
		return err
	}
	if _, err = writer.Write(message); err != nil {
		return err
	}
	if err = writer.Close(); err != nil {
		return err
	}
	return client.Quit()
}
