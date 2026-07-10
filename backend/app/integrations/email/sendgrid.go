package email

import (
	"log"
	"os"
	"path/filepath"

	"backend.app/common/helpers"
	"backend.app/configs"
	"backend.app/internal/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	config *configs.Config
	client *sendgrid.Client
}

func NewSendGrid(config *configs.Config) EmailService {
	sendGrid := &SendGrid{config: config, client: sendgrid.NewSendClient(config.SendGridApiKey)}
	service := EmailService(sendGrid)

	return service

}

func (s *SendGrid) SendEmail(job models.NotificationJob) error {
	m := mail.NewV3Mail()

	body, err := parseEmailTemplate(job)
	if err != nil {
		return err
	}
	//job.From = "trysecured@gmail.com"
	from := mail.NewEmail("", job.From)
	m.SetFrom(from)
	m.AddContent(mail.NewContent("text/html", *body))

	p := mail.NewPersonalization()

	// Add To
	var toAddress []*mail.Email
	for _, email := range job.To {
		toAddress = append(toAddress, mail.NewEmail("", email))
	}

	p.AddTos(toAddress...)

	// Add CCs
	//ccEmails := []string{}
	for _, email := range job.CcEmails {
		//ccEmails = append(ccEmails, email)
		p.AddCCs(mail.NewEmail("", email))
	}
	p.Subject = "Secure Education Email"
	if job.Content["subject"] != nil {
		p.Subject = job.Content["subject"].(string)
	}

	m.AddPersonalizations(p)

	resp, err := s.client.Send(m)
	log.Printf("smtp response: %v", resp)
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	return nil
}

func parseEmailTemplate(job models.NotificationJob) (*string, error) {
	cwd, _ := os.Getwd()
	templateFileName := filepath.Join(cwd, "templates", "email", job.Template+".html")
	return helpers.ParseTemplate(templateFileName, job.Content)
}
