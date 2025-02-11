package sms

import "backend.app/internal/models"

type PhoneService interface {
	Send(job models.NotificationJob, channels string) error
}
