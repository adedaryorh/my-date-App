package core

import (
	"context"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// Logout logs the user out
func (c *Core) Logout(ctx context.Context, user *models.User) error {
	return nil
}

// ChangePassword method used to change an existing password
func (c *Core) ChangePassword(ctx context.Context, data *dtos.ChangePassword, user *models.User) *dtos.ResponseObject {
	// validate old password
	if isValid := helpers.CompareHash(user.PasswordHash, data.OldPassword); !isValid {
		return response.BadRequestResponse(messages.ErrIncorrectPassword)
	}
	// update user with new
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"password_hash": helpers.Hash(data.NewPassword)}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.PasswordChangedSuccessfully, nil)
}

// AddAreaOfInterests methods that adds the user's area of interest
func (c *Core) AddAreaOfInterests(ctx context.Context, data *dtos.AreaOfInterest, user *models.User) *dtos.ResponseObject {
	interests := helpers.RemoveDuplicates(data.AreasOfInterest)
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"interests": interests}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.AreaOfInterestsAddedSuccessfully, nil)
}

// AddPreferredLanguage adds preferred language for the user
func (c *Core) AddPreferredLanguage(ctx context.Context, data *dtos.Language, user *models.User) *dtos.ResponseObject {
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"language": data.Language}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.LanguageAddedSuccessfully, nil)
}

// AddNotificationPreference adds notification preferences for the user
func (c *Core) AddNotificationPreference(ctx context.Context, data *dtos.NotificationPreference, user *models.User) *dtos.ResponseObject {
	preferences := helpers.RemoveDuplicates(data.NotificationPreferences)
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"notification_preference": preferences}); err != nil {
		return response.ServerErrorResponse(err)
	}

	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}
