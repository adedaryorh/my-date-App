package dtos

import (
	"time"
)

type NotificationMetaDataDto struct {
	Action    string `json:"action,omitempty"`
	ObjectRef string `json:"object_ref,omitempty"`
	ObjectId  string `json:"object_id,omitempty"` // Using string to represent UUID in DTO
}

type NotificationDto struct {
	Id               string                  `json:"id"`
	Owner            string                  `json:"owner,omitempty"`
	OwnerId          string                  `json:"owner_id,omitempty"`
	Status           string                  `json:"status,omitempty"`
	Template         string                  `json:"template"`
	Title            string                  `json:"title"`
	NotificationType string                  `json:"notification_type"`
	Content          string                  `json:"content"`
	MetaData         NotificationMetaDataDto `json:"meta_data,omitempty"`
	CreatedAt        time.Time               `json:"created_at,omitempty"`
	UpdatedAt        time.Time               `json:"updated_at,omitempty"`
}

// NotificationsResponse represents a list of notifications.
type NotificationsResponse struct {
	Notifications []NotificationDto `json:"notifications"`
	Total         int               `json:"total"`
	PagingInfo    PagingInfo        `json:"paging_info"`
}
