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

type UserInfo struct {
	UserID       string  `json:"user_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Username     *string `json:"username"`
	BusinessName *string `json:"business_name"`
	IsVerified   bool    `json:"is_verified"`
	ProfileImage *string `json:"profile_image"`
}

type UserProfile struct {
	UserID             string        `json:"user_id"`
	FirstName          string        `json:"first_name"`
	LastName           string        `json:"last_name"`
	Username           string        `json:"username"`
	CountryCode        *string       `json:"country_code"`
	PhoneNumber        *string       `json:"phone_number"`
	Email              string        `json:"email"`
	DateOfBirth        time.Time     `json:"dob"`
	Gender             *string       `json:"gender"`
	RelationshipStatus *string       `json:"relationship_status"`
	Interests          *string       `json:"interests"`
	AccountType        int           `json:"account_type"`
	BusinessInfo       BusinessInfo  `json:"business_info"`
	BasicUserInfo      BasicUserInfo `json:"basic_user_info"`
	ContentsCount      int           `json:"contents_count"`
	CelebrationsCount  int           `json:"celebrations_count"`
	Bio                *string       `json:"bio"`
	IsVerified         bool          `json:"is_verified"`
	ProfileImage       *string       `json:"profile_image"`
	Location           *string       `json:"location"`
}

type BasicUserInfo struct {
	CustomersCount  int `json:"customers_count"`
	LovedOnesCount  int `json:"loved_ones_count"`
	BusinessesCount int `json:"businesses_count"`
}

type BusinessInfo struct {
	CustomersCount int  `json:"customers_count"`
	FollowersCount int  `json:"followers_count"`
	IsFollower     bool `json:"is_followed"`
}
