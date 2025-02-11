package email

import "backend.app/internal/models"

type EmailService interface {
	SendEmail(job models.NotificationJob) error
}
