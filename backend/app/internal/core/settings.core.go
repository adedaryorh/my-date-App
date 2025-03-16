package core

import (
	"context"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
	"github.com/google/uuid"
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

// TogglePushNotification toggles push notification
func (c *Core) TogglePushNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings
	pushSettings.Enabled = !pushSettings.Enabled
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleLikesNotification toggles likes notification
func (c *Core) ToggleLikesNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.Likes = !pushSettings.Likes
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleCommentsNotification toggles comments notification
func (c *Core) ToggleCommentsNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.Comments = !pushSettings.Comments
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleTagsAndMentionsNotification toggles tags and mentions notification
func (c *Core) ToggleTagsAndMentionsNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.TagsAndMentions = !pushSettings.TagsAndMentions
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleRepostNotification toggles repost notification
func (c *Core) ToggleRepostNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.Repost = !pushSettings.Repost
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleDirectMessageNotification toggles direct message notification
func (c *Core) ToggleDirectMessageNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.DirectMessage = !pushSettings.DirectMessage
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleLiveNotification toggles live notification
func (c *Core) ToggleLiveNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.Live = !pushSettings.Live
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// ToggleNewFollowerNotification toggles new follower notification
func (c *Core) ToggleNewFollowerNotification(ctx context.Context, user *models.User) *dtos.ResponseObject {
	pushSettings := user.PushNotificationSettings

	if !pushSettings.Enabled {
		return response.BadRequestResponse(messages.ErrPushNotificationNotEnabled)
	}
	pushSettings.NewFollower = !pushSettings.NewFollower
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"push_notification_settings": pushSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.NotificationPreferenceSetSuccessfully, nil)
}

// SetContentSettings set content settings
func (c *Core) SetContentSettings(ctx context.Context, user *models.User, data *dtos.ContentSettingsDto) *dtos.ResponseObject {
	contentSettings := user.ContentSettings
	contentSettings.Enabled = string(data.Allowed)
	if len(data.SelectedFriends) > 0 && data.Allowed == constants.ContentViewerSelectedFriends {
		var friends []uuid.UUID
		for _, friend := range data.SelectedFriends {
			friendId, _ := uuid.Parse(friend)
			friends = append(friends, friendId)
		}
		contentSettings.SelectedFriends = friends
	}

	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"content_settings": contentSettings}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.ContentSettingsSetSuccessfully, nil)
}

// SetBannedWords used to manage user's banned words
func (c *Core) SetBannedWords(ctx context.Context, data *dtos.SetBannedWordsDto, user *models.User) *dtos.ResponseObject {
	if err := c.repo.UpdateUser(ctx, user.ID, helpers.Map{"banned_words": data.BannedWords}); err != nil {
		return response.ServerErrorResponse(err)
	}
	return response.SuccessResponse(constants.ContentSettingsSetSuccessfully, nil)
}
