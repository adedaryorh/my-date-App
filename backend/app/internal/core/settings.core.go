package core

import (
	"context"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
)

func (c *Core) Logout(ctx context.Context, user *models.User) error {
	return nil
}

// ChangePassword method used to change an existing password
func (c *Core) ChangePassword(ctx context.Context, data *dtos.ChangePassword, user *models.User) *dtos.ResponseObject {
	// validate old password
	if isValid := helpers.CompareHash(user.PasswordHash, data.OldPassword); !isValid {
		return BadRequestResponse(messages.ErrIncorrectPassword)
	}
	// update user with new
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"password_hash": helpers.Hash(data.NewPassword)}); err != nil {
		return ServerErrorResponse(err)
	}
	return SuccessResponse(constants.PasswordChangedSuccessfully, nil)
}

// AddAreaOfInterests
func (c *Core) AddAreaOfInterests(ctx context.Context) {

}

// AddPreferredLanguage
// AddNotificationPreference
