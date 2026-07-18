//go:build ignore

package email

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/sendgrid/sendgrid-go/config"
	"github.com/sendgrid/sendgrid-go/src/models"
	"github.com/rs/zerolog/log"
)

type MailgunSMTP struct {
	config *config.ConfigType
}

func NewMailgun(config *config.ConfigType) EmailService {
	return &MailgunSMTP{
		config: config,
	}
}

func (m *MailgunSMTP) SendEmail(job models.NotificationJob) error {
	log.Info().
		Str("host", m.config.MailHost).
		Str("port", m.config.MailPort).
		Msgf("SendEmail called — port strategy: %s", m.config.MailPort)

	body, err := parseEmailTemplate(job)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse email template")
		return err
	}

	fromAddress := job.From
	if job.FromName != "" {
		fromAddress = fmt.Sprintf("%s <%s>", job.FromName, job.From)
	}

	subject := job.Subject
	if subject == "" && job.Content != nil {
		if s, ok := job.Content["subject"].(string); ok {
			subject = s
		}
	}
	if subject == "" {
		subject = "Notification from PASCA"
	}

	boundary := "pasca-boundary"
	headers := make(map[string]string)
	headers["From"] = fromAddress
	headers["To"] = strings.Join(job.To, ", ")
	if len(job.CcEmails) > 0 {
		headers["Cc"] = strings.Join(job.CcEmails, ", ")
	}
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = fmt.Sprintf("multipart/mixed; boundary=%s", boundary)

	var message strings.Builder
	for k, v := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	message.WriteString("\r\n")

	// HTML body
	message.WriteString(fmt.Sprintf("--%s\r\n", boundary))
	message.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
	message.WriteString(*body)
	message.WriteString("\r\n")

	// Attachments
	for _, att := range job.Attachments {
		if len(att.Data) == 0 || att.Filename == "" {
			continue
		}
		contentType := att.ContentType
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		encoded := base64.StdEncoding.EncodeToString(att.Data)
		message.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		message.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", contentType, att.Filename))
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", att.Filename))
		message.WriteString(encoded)
		message.WriteString("\r\n")
	}
	message.WriteString(fmt.Sprintf("--%s--\r\n", boundary))

	allRecipients := append(job.To, job.CcEmails...)
	smtpAddr := fmt.Sprintf("%s:%s", m.config.MailHost, m.config.MailPort)

	switch m.config.MailPort {
	case "465":
		err = sendWithImplicitTLS(smtpAddr, m.config.MailHost, m.config.MailUsername, m.config.MailPassword, job.From, allRecipients, message.String())
	case "587":
		// STARTTLS
		err = sendWithSTARTTLS(smtpAddr, m.config.MailHost, m.config.MailUsername, m.config.MailPassword, job.From, allRecipients, message.String())
	default:
		// Fallback for Mailtrap
		auth := smtp.PlainAuth("", m.config.MailUsername, m.config.MailPassword, m.config.MailHost)
		err = smtp.SendMail(smtpAddr, auth, job.From, allRecipients, []byte(message.String()))
	}

	// else fallback to HTTP API
	if err != nil {
		log.Warn().Err(err).Msg("SMTP failed, falling back to Resend HTTP API")
		err = sendWithResendAPI(m.config.MailPassword, fromAddress, job.To, job.CcEmails, subject, *body, job.Attachments)
	}

	if err != nil {
		log.Error().Err(err).
			Str("host", m.config.MailHost).
			Str("port", m.config.MailPort).
			Msg("Failed to send email via both SMTP and Resend API")
		return err
	}

	log.Info().Msgf("Email sent successfully - To: %v, Subject: %s", job.To, subject)
	return nil
}

// sendWithSTARTTLS handles port 587.
// Connects on plain TCP first, then upgrades the connection to TLS.
// This is what Resend and Mailgun use on port 587.
func sendWithSTARTTLS(addr, host, username, password, from string, to []string, message string) error {
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return fmt.Errorf("TCP dial failed: %w", err)
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("SMTP client creation failed: %w", err)
	}
	defer client.Close()

	// Upgrade to TLS
	tlsConfig := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: false,
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("STARTTLS failed: %w", err)
	}

	auth := smtp.PlainAuth("", username, password, host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP auth failed: %w", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("SMTP MAIL FROM failed: %w", err)
	}

	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			return fmt.Errorf("SMTP RCPT TO failed for %s: %w", recipient, err)
		}
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("SMTP DATA command failed: %w", err)
	}

	if _, err = writer.Write([]byte(message)); err != nil {
		return fmt.Errorf("SMTP write failed: %w", err)
	}

	if err = writer.Close(); err != nil {
		return fmt.Errorf("SMTP writer close failed: %w", err)
	}

	return client.Quit()
}

func sendWithImplicitTLS(addr, host, username, password, from string, to []string, message string) error {
	tlsConfig := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: false,
	}

	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 10 * time.Second}, "tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("TLS dial failed: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("SMTP client creation failed: %w", err)
	}
	defer client.Close()

	auth := smtp.PlainAuth("", username, password, host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP auth failed: %w", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("SMTP MAIL FROM failed: %w", err)
	}

	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			return fmt.Errorf("SMTP RCPT TO failed for %s: %w", recipient, err)
		}
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("SMTP DATA command failed: %w", err)
	}

	if _, err = writer.Write([]byte(message)); err != nil {
		return fmt.Errorf("SMTP write failed: %w", err)
	}

	if err = writer.Close(); err != nil {
		return fmt.Errorf("SMTP writer close failed: %w", err)
	}

	return client.Quit()
}

func sendWithResendAPI(apiKey, from string, to, cc []string, subject, html string, attachments []models.EmailAttachment) error {
	type resendAttachment struct {
		Filename    string `json:"filename"`
		Content     string `json:"content"`
		ContentType string `json:"content_type,omitempty"`
	}

	type resendPayload struct {
		From        string             `json:"from"`
		To          []string           `json:"to"`
		Cc          []string           `json:"cc,omitempty"`
		Subject     string             `json:"subject"`
		HTML        string             `json:"html"`
		Attachments []resendAttachment `json:"attachments,omitempty"`
	}

	payload := resendPayload{
		From:    from,
		To:      to,
		Subject: subject,
		HTML:    html,
	}

	if len(cc) > 0 {
		payload.Cc = cc
	}

	for _, att := range attachments {
		if len(att.Data) == 0 || att.Filename == "" {
			continue
		}
		contentType := att.ContentType
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		payload.Attachments = append(payload.Attachments, resendAttachment{
			Filename:    att.Filename,
			Content:     base64.StdEncoding.EncodeToString(att.Data),
			ContentType: contentType,
		})
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal resend payload: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create resend request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("resend HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("resend API error: status %d", resp.StatusCode)
	}

	log.Info().Msgf("Email sent via Resend HTTP API")
	return nil
}
