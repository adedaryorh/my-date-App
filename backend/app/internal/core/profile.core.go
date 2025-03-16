package core

import (
	"context"
	"time"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/models"
	"backend.app/internal/services/upload"
	"backend.app/pkg/response"
	"github.com/google/uuid"
)

// UploadUserProfileImage Method used to upload user's profile picture
func (c *Core) UploadUserProfileImage(ctx context.Context, user *models.User, data *dtos.UploadImage) *dtos.ResponseObject {
	maxFileSize := 1024 * 1024 * 10 // 2 MB
	attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG}
	profilePix, err := c.uploadDocument(int64(maxFileSize), string(constants.DocumentKindProfilePicture), user.ID.String(), data.Image, attachmentKinds)
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"profile_image_url": profilePix.Url}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserProfileSuccessfullyUploaded, nil)
}

// UpdateUserProfile Method used to update the user's profile
func (c *Core) UpdateUserProfile(ctx context.Context, user *models.User, data *dtos.UpdateUserProfile) *dtos.ResponseObject {
	// get user
	update := make(helpers.Map)
	if data.Address != nil {
		update["address"] = *data.Address
	}
	if data.FullName != nil {
		firstName, lastName := splitFullName(*data.FullName)
		update["first_name"] = firstName
		if lastName != nil {
			update["last_name"] = *lastName
		}
	}
	if data.DateOfBirth != nil {
		dob, _ := time.Parse(constants.DATE_LAYOUT, *data.DateOfBirth)
		update["date_of_birth"] = dob
	}
	if err := c.repo.UpdateUser(ctx, user.ID, update); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserProfileSuccessfullyUpdated, nil)
}

// GetSelfProfile get self profile
func (c *Core) GetSelfProfile(ctx context.Context, user *models.User) *dtos.ResponseObject {
	// map user
	var profile mappers.DtoUserMapper
	return response.SuccessResponse(constants.UserSuccessFullyFetched, profile.MapSelfProfileDto(user))
}

// GetSingleProfile get single user profile
func (c *Core) GetSingleProfile(ctx context.Context, userId uuid.UUID, user *models.User) *dtos.ResponseObject {
	var profile mappers.DtoUserMapper
	// TODO !!!! check if user blocked

	// get user
	userData, err := c.repo.GetUserByField(ctx, helpers.Map{"id": userId})
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	// map user
	return response.SuccessResponse(constants.UserSuccessFullyFetched, profile.MapUserProfileDto(userData))
}

// GetAllProfiles get all user profiles
func (c *Core) GetAllProfiles(ctx context.Context, user *models.User, query *dtos.APIPagingDto) *dtos.ResponseObject {
	result, err := c.repo.GetAllUsers(ctx, user, query)
	if err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.UserProfilesSuccessfullyFetched, result)
}
