package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// User this is the user's model object
type User struct {
	ID                 uuid.UUID       `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	FirstName          string          `json:"first_name,omitempty"`
	LastName           *string         `json:"last_name,omitempty"`
	BusinessName       *string         `json:"business_name,omitempty"`
	IndustryType       *string         `json:"industry_type,omitempty"`
	Username           string          `json:"username,omitempty"`
	CountryCode        string          `json:"country_code,omitempty"`
	Latitude           float64         `json:"latitude,omitempty"`
	Longitude          float64         `json:"longitude,omitempty"`
	PhoneNumber        string          `json:"phone_number,omitempty"`
	Email              string          `json:"email,omitempty"`
	CompletionState    int             `json:"completion_state,omitempty"`
	IpAddress          *string         `json:"ip_address,omitempty"`
	DeviceType         *string         `json:"device_type,omitempty"`
	DateOfBirth        *time.Time      `json:"date_of_birth,omitempty"`
	AccountType        string          `json:"account_type,omitempty"`
	Interests          *pq.StringArray `json:"interests,omitempty"`
	ProfileImageURL    *string         `json:"profile_image_url,omitempty"`
	PasswordHash       string          `json:"-"`
	Status             string          `json:"status,omitempty"`
	VerificationStatus string          `json:"verification_status,omitempty"`
	NextLoginAt        *time.Time      `json:"next_login_at,omitempty"`
	CreatedAt          time.Time       `json:"created_at,omitempty"`
	UpdatedAt          *time.Time      `json:"updated_at,omitempty"`
}

// AWSObjectUrl this is the uploaded file object
type AWSObjectUrl struct {
	KeyName   string     `json:"key_name,omitempty"`
	Url       string     `json:"url,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Note      string     `json:"note,omitempty"`
}
