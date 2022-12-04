package models

import (
	"mime/multipart"
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
	ProfileImage       *multipart.FileHeader
	Password           string
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
