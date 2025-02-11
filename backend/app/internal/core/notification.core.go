package core

import (
	"context"
	"errors"
	"fmt"

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
