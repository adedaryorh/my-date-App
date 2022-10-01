package dtos

import (
	"time"
)

type User struct {
	UserID             string    `json:"user_id"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	Username           string    `json:"username"`
	CountryCode        *string   `json:"country_code"`
	PhoneNumber        *string   `json:"phone_number"`
	Email              string    `json:"email"`
	DateOfBirth        time.Time `json:"dob"`
	Gender             *string   `json:"gender"`
	RelationshipStatus *string   `json:"relationship_status"`
	Interests          *string   `json:"interests"`
}
