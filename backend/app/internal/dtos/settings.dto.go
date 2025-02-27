package dtos

import "backend.app/common/constants"

// AreaOfInterest  the dto for setting user's area of interest
type AreaOfInterest struct {
	AreasOfInterest []constants.AreaOfInterest `json:"areas_of_interest" validate:"gt=0,dive,is_enum"`
}

// Language  the dto for setting user's language
type Language struct {
	Language string `json:"language" validate:"required,min=2,max=100"`
}

// NotificationPreference the dto for setting user's notification preference
type NotificationPreference struct {
	NotificationPreferences []constants.NotificationType `json:"notification_preferences" validate:"gt=0,dive,is_enum"`
}
