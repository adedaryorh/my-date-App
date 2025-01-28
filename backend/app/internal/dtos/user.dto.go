package dtos

import (
	"time"

	"backend.app/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	UserID             uuid.UUID      `json:"user_id"`
	FirstName          string         `json:"first_name"`
	LastName           string         `json:"last_name"`
	Username           string         `json:"username"`
	CountryCode        string         `json:"country_code"`
	PhoneNumber        string         `json:"phone_number"`
	Email              string         `json:"email"`
	DateOfBirth        *time.Time     `json:"dob"`
	Gender             *string        `json:"gender"`
	RelationshipStatus *string        `json:"relationship_status"`
	Interests          pq.StringArray `json:"interests"`
	ProfileImage       *string        `json:"profile_image"`
	DateJoined         time.Time      `json:"date_joined"`
}

type UserInfo struct {
	UserID       uuid.UUID `json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	IsVerified   bool      `json:"is_verified"`
	ProfileImage *string   `json:"profile_image"`
}

type UserProfile struct {
	UserID             uuid.UUID      `json:"user_id"`
	FirstName          string         `json:"first_name"`
	LastName           string         `json:"last_name"`
	Username           string         `json:"username"`
	CountryCode        *string        `json:"country_code"`
	PhoneNumber        *string        `json:"phone_number"`
	Email              string         `json:"email"`
	DateOfBirth        time.Time      `json:"dob"`
	Gender             *string        `json:"gender"`
	RelationshipStatus *string        `json:"relationship_status"`
	Interests          pq.StringArray `json:"interests"`
	AccountType        AccountType    `json:"account_type"`
	ContentsCount      int            `json:"contents_count"`
	CelebrationsCount  int            `json:"celebrations_count"`
	LovedOnesCount     int            `json:"loved_ones_count"`
	BusinessesCount    int            `json:"businesses_count"`
	CustomersCount     int            `json:"customers_count"`
	Bio                *string        `json:"bio"`
	IsVerified         bool           `json:"is_verified"`
	ProfileImage       *string        `json:"profile_image"`
	Location           *string        `json:"location"`
}

type AccountType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type IndustryType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UsersResponse struct {
	Users      []*models.User `json:"users"`
	PagingInfo PagingInfo     `json:"paging_info"`
}
