package mappers

import (
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

type UserMapper interface {
	//MapToUserDto(user models.User) dtos.User
	MapToUserProfileDto(user models.User) dtos.UserProfile
	MapToUserInfoDto(user models.User) dtos.UserInfo
}

// type DtoUserMapper struct {
// }

// func (um *DtoUserMapper) MapToUserDto(user models.User) dtos.User {
// 	dto := dtos.User{
// 		ID:          user.ID,
// 		FirstName:   user.FirstName,
// 		LastName:    user.LastName,
// 		Username:    user.Username,
// 		CountryCode: user.CountryCode,
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		DateOfBirth: &user.DateOfBirth,

// 		Interests:    user.Interests,
// 		ProfileImage: user.ProfileImageURL,
// 		DateJoined:   user.CreatedAt,
// 	}

// 	// if user.AccountType.ID == config.ACCOUNT_BASIC_ID {
// 	// 	dto.AccountType.Name = config.ACCOUNT_BASIC
// 	// } else {
// 	// 	dto.AccountType.Name = config.ACCOUNT_BUSINESS
// 	// }

// 	return dto
// }

// func (um *DtoUserMapper) MapToUserProfileDto(user models.User) dtos.UserProfile {
// 	return dtos.UserProfile{
// 		UserID:       user.ID,
// 		FirstName:    user.FirstName,
// 		LastName:     user.LastName,
// 		Username:     user.Username,
// 		CountryCode:  &user.CountryCode,
// 		PhoneNumber:  &user.PhoneNumber,
// 		Email:        user.Email,
// 		DateOfBirth:  user.DateOfBirth,
// 		Interests:    user.Interests,
// 		ProfileImage: user.ProfileImageURL,
// 	}
// }

// func (um *DtoUserMapper) MapToUserInfoDto(user models.User) dtos.UserInfo {
// 	return dtos.UserInfo{
// 		UserID:       user.ID,
// 		FirstName:    user.FirstName,
// 		LastName:     user.LastName,
// 		Username:     user.Username,
// 		ProfileImage: user.ProfileImageURL,
// 	}
// }
