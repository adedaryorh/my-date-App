package dtos

import (
	"mime/multipart"
	"time"

	"backend.app/common/constants"
	"backend.app/internal/models"
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
