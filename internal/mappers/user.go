package mappers

import (
	"celebut-api/internal/config"
	"celebut-api/internal/dtos"
	"celebut-api/internal/models"
)

type UserMapper interface {
	MapToUserDto(user models.User) dtos.User
	MapToUserProfileDto(user models.User) dtos.UserProfile
	MapToUserInfoDto(user models.User) dtos.UserInfo
}

type DtoUserMapper struct {
}

func (um *DtoUserMapper) MapToUserDto(user models.User) dtos.User {
	dto := dtos.User{
		UserID:             user.UserID,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		Username:           user.Username,
		CountryCode:        user.CountryCode,
		PhoneNumber:        user.PhoneNumber,
		Email:              user.Email,
		DateOfBirth:        user.DateOfBirth,
		Gender:             user.Gender,
		RelationshipStatus: user.RelationshipStatus,
		Interests:          user.Interests,
		BusinessName:       user.BusinessName,
		ProfileImage:       user.ProfileImageURL,
		AccountType: dtos.AccountType{
			ID: user.AccountType.ID,
		},
	}

	if user.AccountType.ID == config.ACCOUNT_BASIC_ID {
		dto.AccountType.Name = config.ACCOUNT_BASIC
	} else {
		dto.AccountType.Name = config.ACCOUNT_BUSINESS
	}

	if user.Industry != nil {
		dto.IndustryType = &dtos.IndustryType{
			ID:   user.Industry.ID,
			Name: user.Industry.Name,
		}
	}

	return dto
}

func (um *DtoUserMapper) MapToUserProfileDto(user models.User) dtos.UserProfile {
	return dtos.UserProfile{
		UserID:             user.UserID,
		FirstName:          *user.FirstName,
		LastName:           *user.LastName,
		Username:           *user.Username,
		CountryCode:        user.CountryCode,
		PhoneNumber:        user.PhoneNumber,
		Email:              user.Email,
		DateOfBirth:        *user.DateOfBirth,
		Gender:             user.Gender,
		RelationshipStatus: user.RelationshipStatus,
		Interests:          user.Interests,
		ProfileImage:       user.ProfileImageURL,
		AccountType: dtos.AccountType{
			ID: user.AccountType.ID,
		},
	}
}

func (um *DtoUserMapper) MapToUserInfoDto(user models.User) dtos.UserInfo {
	return dtos.UserInfo{
		UserID:       user.UserID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		ProfileImage: user.ProfileImageURL,
	}
}
