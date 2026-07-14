package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

// User this is the user's model object
type User struct {
	ID                       uuid.UUID            `json:"id" gorm:"column:id;PRIMARY_KEY;type:uuid;default:gen_random_uuid()"`
	FirstName                string               `json:"first_name,omitempty"`
	LastName                 *string              `json:"last_name,omitempty"`
	BusinessName             *string              `json:"business_name,omitempty"`
	IndustryType             *string              `json:"industry_type,omitempty"`
	Username                 string               `json:"username,omitempty"`
	CountryCode              string               `json:"country_code,omitempty"`
	Latitude                 float64              `json:"latitude,omitempty"`
	Longitude                float64              `json:"longitude,omitempty"`
	PhoneNumber              string               `json:"phone_number,omitempty"`
	Email                    string               `json:"email,omitempty"`
	CompletionState          int                  `json:"completion_state,omitempty"`
	IpAddress                *string              `json:"ip_address,omitempty"`
	DeviceType               *string              `json:"device_type,omitempty"`
	DateOfBirth              *time.Time           `json:"date_of_birth,omitempty"`
	AccountType              string               `json:"account_type,omitempty"`
	Interests                *pq.StringArray      `json:"interests,omitempty"`
	ProfileImageURL          *string              `json:"profile_image_url,omitempty"`
	Language                 *string              `json:"language,omitempty"`
	NotificationPreference   *pq.StringArray      `json:"notification_preference,omitempty"`
	PasswordHash             string               `json:"-"`
	Status                   string               `json:"status,omitempty"`
	VerificationStatus       string               `json:"verification_status,omitempty"`
	NextLoginAt              *time.Time           `json:"next_login_at,omitempty"`
	Followers                int64                `json:"followers,omitempty"`
	Following                int64                `json:"following,omitempty"`
	Blocked                  int64                `json:"blocked,omitempty"`
	PushNotificationSettings PushNotificationData `json:"push_notification_settings,omitempty"`
	ContentSettings          ContentSettings      `json:"content_settings,omitempty"`
	BannedWords              *pq.StringArray      `json:"banned_words"`
	Role                     string               `json:"role,omitempty" gorm:"default:'user'"`
	GoogleID                 string               `json:"google_id,omitempty" gorm:"column:google_id"`
	CreatedAt                time.Time            `json:"created_at,omitempty"`
	UpdatedAt                *time.Time           `json:"updated_at,omitempty"`
}

type PushNotificationData struct {
	Enabled         bool `json:"push_enabled"`
	Likes           bool `json:"likes"`
	Comments        bool `json:"comments"`
	TagsAndMentions bool `json:"tags_and_mentions"`
	Repost          bool `json:"repost"`
	DirectMessage   bool `json:"direct_message"`
	Live            bool `json:"live"`
	NewFollower     bool `json:"new_follower"`
}

type ContentSettings struct {
	Enabled         string      `json:"enabled"`
	SelectedFriends []uuid.UUID `json:"selected_friends"`
}

type AWSObjectUrl struct {
	KeyName   string     `json:"key_name,omitempty"`
	Url       string     `json:"url,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Note      string     `json:"note,omitempty"`
}

type Incrementor struct {
	Field    string
	Operator string
	Value    int64
}

func (p PushNotificationData) Value() (driver.Value, error) { return json.Marshal(p) }
func (p *PushNotificationData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, p)
}
func (c ContentSettings) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *ContentSettings) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, c)
}
