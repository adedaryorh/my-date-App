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

	c.log.Info("Attempting to send notification",
		zap.String("user_id", user.ID.String()),
		zap.String("user_email", user.Email),
		zap.String("template", template),
		zap.String("action", "send_notification"))

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
		c.log.Error("Failed to send notification",
			zap.String("user_id", user.ID.String()),
			zap.String("template", template),
			zap.String("channels", fmt.Sprintf("%v", channels)),
			zap.Error(fmt.Errorf(errString)),
			zap.String("action", "send_notification"))
		return errors.New(errString)
	}

	c.log.Info("Successfully sent notification",
		zap.String("user_id", user.ID.String()),
		zap.String("template", template),
		zap.String("channels", fmt.Sprintf("%v", channels)),
		zap.String("action", "send_notification"))
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
	c.log.Info("Creating notification",
		zap.String("user_id", user.ID.String()),
		zap.String("user_email", user.Email),
		zap.String("template", template),
		zap.String("notification_type", notificationType),
		zap.String("action", "create_new_notification"))

	newNotification, err := c.repo.CreateNotification(ctx, notification)
	if err != nil {
		c.log.Error("Failed to create notification",
			zap.String("user_id", user.ID.String()),
			zap.String("template", template),
			zap.Error(err),
			zap.String("action", "create_new_notification"))
		return nil, err
	}

	c.log.Info("Successfully created notification",
		zap.String("notification_id", newNotification.Id.String()),
		zap.String("user_id", user.ID.String()),
		zap.String("template", template),
		zap.String("action", "create_new_notification"))
	return newNotification, nil
}

// CreateAndSendNotification
// GetAllNotifications returns all notifications with pagination
func (c *Core) GetAllNotifications(ctx context.Context, query *dtos.APIPagingDto) *dtos.ResponseObject {
	c.log.Info("Retrieving notifications list",
		zap.Int("limit", query.Limit),
		zap.Int("page", query.Page),
		zap.String("action", "get_notifications"))

	result, err := c.repo.GetAllNotifications(ctx, query)
	if err != nil {
		c.log.Error("Failed to retrieve notifications",
			zap.Error(err),
			zap.String("action", "get_notifications"))
		return response.ServerErrorResponse(err)
	}

	c.log.Info("Successfully retrieved notifications",
		zap.Int("count", len(result.Notifications)),
		zap.String("action", "get_notifications"))
	return response.SuccessResponse("notifications retrieved successfully", result)
}

// GetNotificationById returns a notification by its ID
func (c *Core) GetNotificationById(ctx context.Context, notificationId uuid.UUID) *dtos.ResponseObject {
	c.log.Info("Retrieving notification by ID",
		zap.String("notification_id", notificationId.String()),
		zap.String("action", "get_notification"))

	result, err := c.repo.GetNotificationById(ctx, notificationId)
	if err != nil {
		c.log.Error("Failed to retrieve notification",
			zap.String("notification_id", notificationId.String()),
			zap.Error(err),
			zap.String("action", "get_notification"))
		return response.ServerErrorResponse(err)
	}
	if result == nil {
		c.log.Warn("Notification not found",
			zap.String("notification_id", notificationId.String()),
			zap.String("action", "get_notification"))
		return response.NotFoundResponse(fmt.Errorf("notification not found"), "Notification not found")
	}

	c.log.Info("Successfully retrieved notification",
		zap.String("notification_id", notificationId.String()),
		zap.String("action", "get_notification"))
	return response.SuccessResponse("notification retrieved successfully", result)
}

// DeleteNotification deletes a notification by its ID
func (c *Core) DeleteNotification(ctx context.Context, notificationId uuid.UUID) *dtos.ResponseObject {
	c.log.Info("Attempting to delete notification",
		zap.String("notification_id", notificationId.String()),
		zap.String("action", "delete_notification"))

	// First check if notification exists
	existing, _ := c.repo.GetNotificationById(ctx, notificationId)
	if existing == nil {
		c.log.Warn("Attempt to delete non-existent notification",
			zap.String("notification_id", notificationId.String()),
			zap.String("action", "delete_notification"))
		return response.NotFoundResponse(fmt.Errorf("notification not found"), "Notification not found")
	}

	err := c.repo.DeleteNotification(ctx, notificationId)
	if err != nil {
		c.log.Error("Failed to delete notification",
			zap.String("notification_id", notificationId.String()),
			zap.Error(err),
			zap.String("action", "delete_notification"))
		return response.ServerErrorResponse(err)
	}

	c.log.Info("Successfully deleted notification",
		zap.String("notification_id", notificationId.String()),
		zap.String("action", "delete_notification"))
	return response.SuccessResponse("notification deleted successfully", nil)
}

// MarkNotificationAsRead marks a notification as read by updating its status
func (c *Core) MarkNotificationAsRead(ctx context.Context, notificationId uuid.UUID) *dtos.ResponseObject {
	c.log.Info("Attempting to mark notification as read",
		zap.String("notification_id", notificationId.String()),
		zap.String("action", "mark_notification_read"))

	// First check if notification exists
	existing, _ := c.repo.GetNotificationById(ctx, notificationId)
	if existing == nil {
		c.log.Warn("Attempt to mark non-existent notification as read",
			zap.String("notification_id", notificationId.String()),
			zap.String("action", "mark_notification_read"))
		return response.NotFoundResponse(fmt.Errorf("notification not found"), "Notification not found")
	}

	// Update the notification status to read
	fields := map[string]interface{}{"status": models.NotificationStatusRead}
	err := c.repo.UpdateNotification(ctx, notificationId, fields)
	if err != nil {
		c.log.Error("Failed to mark notification as read",
			zap.String("notification_id", notificationId.String()),
			zap.Error(err),
			zap.String("action", "mark_notification_read"))
		return response.ServerErrorResponse(err)
	}

	c.log.Info("Successfully marked notification as read",
		zap.String("notification_id", notificationId.String()),
		zap.String("action", "mark_notification_read"))
	return response.SuccessResponse("notification marked as read successfully", nil)
}
