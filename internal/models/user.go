package models

import (
	"time"
)

type User struct {
	ID                 int
	UserID             string
	FirstName          *string
	LastName           *string
	Username           *string
	CountryCode        *string
	PhoneNumber        *string
	Email              string
	DateOfBirth        *time.Time
	Gender             *string
	RelationshipStatus *string
	BusinessName       *string
	Industry           *Industry
	AccountType        AccountType
	Interests          *string
	ProfileImageURL    *string
	Password           string
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
