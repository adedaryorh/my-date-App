package mappers

import (
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

type DtoUserMapper struct{}

func (um *DtoUserMapper) MapSelfProfileDto(user *models.User) *dtos.UserProfile {
	profile := dtos.UserProfile{
		ID:                     user.ID,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Username:               user.Username,
		CountryCode:            user.CountryCode,
		PhoneNumber:            user.PhoneNumber,
		Email:                  user.Email,
		DateOfBirth:            user.DateOfBirth,
		Interests:              user.Interests,
		ProfileImageUrl:        user.ProfileImageURL,
		Language:               user.Language,
		NotificationPreference: user.NotificationPreference,
		AccountType:            user.AccountType,
		Followers:              helpers.Int64ToPointer(user.Followers),
		Following:              helpers.Int64ToPointer(user.Following),
		Blocked:                helpers.Int64ToPointer(user.Blocked),
		DateJoined:             user.CreatedAt,
	}

	return &profile
}

func (um *DtoUserMapper) MapUserProfileDto(user *models.User) *dtos.UserProfile {
	profile := dtos.UserProfile{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Username:        user.Username,
		DateOfBirth:     user.DateOfBirth,
		Interests:       user.Interests,
		ProfileImageUrl: user.ProfileImageURL,
		AccountType:     user.AccountType,
		Followers:       helpers.Int64ToPointer(user.Followers),
		Following:       helpers.Int64ToPointer(user.Following),
		Blocked:         helpers.Int64ToPointer(user.Blocked),
		DateJoined:      user.CreatedAt,
	}

	return &profile
}

func (um *DtoUserMapper) MapBlockUserProfileDto(user *models.User) *dtos.UserProfile {
	profile := dtos.UserProfile{
		ID:         user.ID,
		BlockedYou: true,
		DateJoined: user.CreatedAt,
	}

	return &profile
}
