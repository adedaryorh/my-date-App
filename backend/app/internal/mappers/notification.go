package mappers

import (
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type DtoNotificationMapper struct{}

// Helper function to map Notification model to DTO.
func (nm *DtoNotificationMapper) MapNotificationToDto(notification *models.Notification) *dtos.NotificationDto {
	return &dtos.NotificationDto{
		Id:               notification.Id.String(),
		Owner:            notification.Owner,
		OwnerId:          notification.OwnerId.String(),
		Status:           string(notification.Status),
		Template:         notification.Template,
		Title:            notification.Title,
		NotificationType: notification.NotificationType,
		Content:          notification.Content,
		MetaData: dtos.NotificationMetaDataDto{
			Action:    notification.MetaData.Action,
			ObjectRef: notification.MetaData.ObjectRef,
			ObjectId:  notification.MetaData.ObjectId.String(),
		},
		CreatedAt: notification.CreatedAt,
		UpdatedAt: notification.UpdatedAt,
	}
}

// method to map a DTO back to a model
func (nm *DtoNotificationMapper) MapDtoToNotification(dto *dtos.NotificationDto) (*models.Notification, error) {
	ownerId, err := uuid.Parse(dto.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("invalid owner ID format: %w", err)
	}

	var objectId uuid.UUID
	if dto.MetaData.ObjectId != "" {
		objectId, err = uuid.Parse(dto.MetaData.ObjectId)
		if err != nil {
			return nil, fmt.Errorf("invalid object ID format: %w", err)
		}
	}

	notificationId, err := uuid.Parse(dto.Id)
	if err != nil && dto.Id != "" {
		return nil, fmt.Errorf("invalid notification ID format: %w", err)
	}

	return &models.Notification{
		Id:               notificationId,
		Owner:            dto.Owner,
		OwnerId:          ownerId,
		Status:           models.NotificationStatus(dto.Status),
		Template:         dto.Template,
		Title:            dto.Title,
		NotificationType: dto.NotificationType,
		Content:          dto.Content,
		MetaData: models.NotificationMetaData{
			Action:    dto.MetaData.Action,
			ObjectRef: dto.MetaData.ObjectRef,
			ObjectId:  objectId,
		},
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}, nil
}
