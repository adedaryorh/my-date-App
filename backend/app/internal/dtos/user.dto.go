package dtos

import (
	"mime/multipart"
	"time"

	"backend.app/common/constants"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// UserSignUp data object to sign up user
type UserSignUp struct {
	FullName    string `json:"full_name" validate:"required,min=3,max=100"`
	Username    string `json:"username" validate:"required,min=3,max=50"`
	CountryCode string `json:"country_code" validate:"required,min=2,max=3"`
	PhoneNumber string `json:"phone_number" validate:"required,is_phone"`
	Password    string `json:"password" validate:"required,is_password"`
	Email       string `json:"email" validate:"required,email"`
	DateOfBirth string `json:"date_of_birth" validate:"required,is_date,before_now"`
}

// BusinessSignUp data object to sign up business
type BusinessSignUp struct {
	BusinessName string                 `json:"business_name" validate:"required,min=3,max=200"`
	Username     string                 `json:"username" validate:"required,min=3,max=50"`
	CountryCode  string                 `json:"country_code" validate:"required,min=2,max=3"`
	PhoneNumber  string                 `json:"phone_number" validate:"required,is_phone"`
	Password     string                 `json:"password" validate:"required,is_password"`
	Email        string                 `json:"email" validate:"required,email"`
	Industry     constants.IndustryType `json:"industry" validate:"required,is_enum"`
}

// ConfirmPhoneNumber data object for phone number confirmation
type ConfirmPhoneNumber struct {
	PhoneNumber string `json:"phone_number" validate:"required,is_phone"`
	Token       string `json:"token" validate:"required,len=6"`
}

// UpdateUserProfile data object to update the user profile
type UpdateUserProfile struct {
	FullName    *string `json:"full_name" validate:"omitempty,min=3,max=100"`
	DateOfBirth *string `json:"date_of_birth" validate:"omitempty,is_date,before_now"`
	Address     *string `json:"address" validate:"omitempty,min=3,max=100"`
}

// Email The email object
type Email struct {
	Email string `json:"email" validate:"required,email"`
}

// Phone The phone object
type Phone struct {
	Phone string `json:"phone_number" validate:"required,is_phone"`
}

// UploadImage data object to update an image
type UploadImage struct {
	Image *multipart.FileHeader `form:"image" validate:"required,is_file"`
}

// ResetPassword data to reset password
type ResetPassword struct {
	PhoneNumber string `json:"phone_number" validate:"required,is_phone"`
	Password    string `json:"password" validate:"required,is_password"`
	Token       string `json:"token" validate:"required,len=6"`
}

// ChangePassword data to change an existing password
type ChangePassword struct {
	OldPassword string `json:"old_password" validate:"required,is_password"`
	NewPassword string `json:"new_password" validate:"required,is_password"`
}

type UserInfo struct {
	UserID       uuid.UUID `json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	IsVerified   bool      `json:"is_verified"`
	ProfileImage *string   `json:"profile_image"`
}

// UserProfile is the user profile object
type UserProfile struct {
	ID                     uuid.UUID       `json:"id"`
	FirstName              string          `json:"first_name"`
	LastName               *string         `json:"last_name,omitempty"`
	Username               string          `json:"username"`
	CountryCode            string          `json:"country_code,omitempty"`
	PhoneNumber            string          `json:"phone_number,omitempty"`
	Email                  string          `json:"email,omitempty"`
	DateOfBirth            *time.Time      `json:"date_of_birth,omitempty"`
	Gender                 *string         `json:"gender,omitempty"`
	RelationshipStatus     *string         `json:"relationship_status"`
	Interests              *pq.StringArray `json:"interests,omitempty"`
	Language               *string         `json:"language,omitempty"`
	NotificationPreference *pq.StringArray `json:"notification_preference,omitempty"`
	AccountType            string          `json:"account_type,omitempty"`
	ContentsCount          int64           `json:"contents_count,omitempty"`
	CelebrationsCount      int64           `json:"celebrations_count,omitempty"`
	LovedOnesCount         *int64          `json:"loved_ones_count,omitempty"`
	BusinessesCount        *int64          `json:"businesses_count,omitempty"`
	CustomersCount         *int64          `json:"customers_count,omitempty"`
	Bio                    *string         `json:"bio,omitempty"`
	VerificationStatus     string          `json:"verification_status,omitempty"`
	ProfileImageUrl        *string         `json:"profile_image_url,omitempty"`
	Location               *string         `json:"location,omitempty"`
	Followers              *int64          `json:"followers,omitempty"`
	Following              *int64          `json:"following,omitempty"`
	Blocked                *int64          `json:"blocked,omitempty"`
	BlockedYou             bool            `json:"blocked_you"`
	DateJoined             time.Time       `json:"date_joined"`
}

// UsersResponse the user's response object with pagination
type UsersResponse struct {
	Users      []*UserProfile `json:"users"`
	PagingInfo PagingInfo     `json:"paging_info"`
}
