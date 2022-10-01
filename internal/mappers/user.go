package mappers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/models"
)

type UserMapper interface {
	MapToUserDto(user models.User) dtos.User
	MapToBusinessDto(user models.User) dtos.Business
}

type DtoUserMapper struct {
}

func (um *DtoUserMapper) MapToUserDto(user models.User) dtos.User {
	return dtos.User{
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
	}
}

func (um *DtoUserMapper) MapToBusinessDto(user models.User) dtos.Business {
	return dtos.Business{
		UserID:       user.UserID,
		BusinessName: *user.BusinessName,
		//IndustryType: user.IndustryId,
		CountryCode: user.CountryCode,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
}
