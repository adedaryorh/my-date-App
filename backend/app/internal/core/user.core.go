package core

import (
	"context"
	"time"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/internal/services/upload"
)

// UploadUserProfileImage Method used to upload user's profile picture
func (c *Core) UploadUserProfileImage(ctx context.Context, user *models.User, data *dtos.UploadImage) *dtos.ResponseObject {
	maxFileSize := 1024 * 1024 * 2 // 2 MB
	attachmentKinds := []upload.AttachmentKind{upload.AttachmentKindImagePNG, upload.AttachmentKindImageJPEG, upload.AttachmentKindImageJPG}
	profilePix, err := c.uploadDocument(int64(maxFileSize), string(constants.DocumentKindProfilePicture), user.ID.String(), data.Image, attachmentKinds)
	if err != nil {
		return ServerErrorResponse(err)
	}
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"profile_image_url": profilePix.Url}); err != nil {
		return ServerErrorResponse(err)
	}
	return SuccessResponse(constants.UserProfileSuccessfullyUploaded, nil)
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
		return ServerErrorResponse(err)
	}
	return SuccessResponse(constants.UserProfileSuccessfullyUpdated, nil)
}
