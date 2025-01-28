package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID                 int
	UserID             uuid.UUID
	FirstName          string
	LastName           string
	Username           string
	CountryCode        string
	PhoneNumber        string
	Email              string
	DateOfBirth        *time.Time
	Gender             *string
	RelationshipStatus *string
	Industry           *Industry
	AccountType        string
	Interests          pq.StringArray
	ProfileImageURL    *string
	PasswordHash       string
	Status             string
	VerificationStatus string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type AWSObjectUrl struct {
	KeyName   string     `json:"keyName,omitempty"`
	Url       string     `json:"url,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Note      string     `json:"note,omitempty"`
}
