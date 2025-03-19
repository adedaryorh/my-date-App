package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"

	"backend.app/internal/models"
)

func (c *Core) SendNotification(ctx context.Context, user *models.User, template string, content map[string]interface{}, opts ...models.NotificationOpts) error {
	var errString string
	var err error
	// extract channels from template
	channels := models.NotificationTemplates[template]

	for _, channel := range channels {
		// if channel == models.NotificationChannels.Email {
		// 	err = c.sendEmail(user, template, content)
		// 	if err != nil {
		// 		errString = err.Error()
		// 	}
		// }

		if channel == models.NotificationChannels.SMS {
			err = c.sendToPhone(user, template, content, channel)
			if err != nil {
				errString = fmt.Sprintf("%s \n %s", errString, err.Error())
			}
		}
		if channel == models.NotificationChannels.Whatsapp {
			err = c.sendToPhone(user, template, content, channel)
			if err != nil {
				errString = fmt.Sprintf("%s \n %s", errString, err.Error())
			}
		}

		// if channel == models.NotificationChannels.Push {
		// 	err = c.sendPush(user, template, content)
		// 	if err != nil {
		// 		errString = fmt.Sprintf("%s \n %s", errString, err.Error())
		// 	}
		// }

	}
	if errString != "" {
		return errors.New(errString)
	}
	return nil
}

// func (c *Core) sendEmail(user *models.User, template string, content map[string]interface{}) error {
// 	// provider would later be set to be dynamic
// 	provider := "sendgrid"

//		// build notification job
//		job := models.NotificationJob{
//			From:     "info@secur.education",
//			To:       []string{user.Email},
//			Template: template,
//			Content:  content,
//		}
//		return c.emailService[provider].SendEmail(job)
//	}
func (c *Core) sendToPhone(user *models.User, template string, content map[string]interface{}, channel string) error {
	// provider would later be set to be dynamic
	provider := "twilio"
	// build notification job
	job := models.NotificationJob{
		To:       []string{user.PhoneNumber},
		Template: template,
		Content:  content,
		From:     "+447380337307",
	}
	err := c.phoneService[provider].Send(job, channel)
	if err != nil {
		return err
	}
	return nil
}

// createNotification
func (c *Core) CreateNotification(ctx context.Context, user *models.User, template string, content map[string]interface{}, opts ...models.NotificationOpts) (*models.Notification, error) {
	var opt models.NotificationOpts
	if len(opts) > 0 {
		opt = opts[0]
	}
	title := opt.Title
	if title == "" {
		title = string(models.NotificationTitleScholarshipUpdate)
	}
	message := opt.Message
	if message == "" && content != nil {
		if msg, ok := content["message"].(string); ok {
			message = msg
		}
	}
	notificationType := opt.NotificationType
	if notificationType == "" {
		notificationType = template
	}
	notification := &models.Notification{
		Id:               uuid.New(),
		Owner:            user.Username,
		OwnerId:          user.ID,
		Status:           models.NotificationStatusSent,
		Template:         template,
		Title:            title,
		NotificationType: notificationType,
		Content:          message,
		MetaData: models.NotificationMetaData{
			Action:    opt.Action,
			ObjectRef: opt.ObjectRef,
			ObjectId:  opt.ObjectId,
		},
	}
	newNotification, err := c.repo.CreateNotification(ctx, notification)
	if err != nil {
		return nil, err
	}
	return newNotification, nil
}

// CreateAndSendNotification
func (c *Core) CreateAndSendNotification(ctx context.Context, user *models.User, template string, content map[string]interface{}, opts ...models.NotificationOpts) (*models.Notification, error) {
	err := c.SendNotification(ctx, user, template, content, opts...)
	if err != nil {
		return nil, err
	}
	notification, err := c.CreateNotification(ctx, user, template, content, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification record: %w", err)
	}
	return notification, nil
}
