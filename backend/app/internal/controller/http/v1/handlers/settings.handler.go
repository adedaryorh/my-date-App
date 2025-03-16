package handlers

import (
	"github.com/gin-gonic/gin"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// @Tags Settings
// @Summary Change user password
// @Description Changes user password
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.ChangePassword  true  "data to change user password "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/change-password [patch]
func (h *Handler) ChangePassword(c *gin.Context) {
	var input *dtos.ChangePassword
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.ChangePassword(c, input, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Add Area Of Interest
// @Description Adds User Area of Interest
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.AreaOfInterest  true  "data to add area of interests "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/area-of-interests [patch]
func (h *Handler) AddAreaOfInterests(c *gin.Context) {
	var input *dtos.AreaOfInterest
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.AddAreaOfInterests(c, input, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Add Preferred Language
// @Description Adds user's Preferred Language
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.Language  true  "data to add preferred language"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/add-language [patch]
func (h *Handler) AddPreferredLanguage(c *gin.Context) {
	var input *dtos.Language
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.AddPreferredLanguage(c, input, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Add Notification Preference
// @Description Adds user's Notification Preference
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.NotificationPreference  true  "data to add notification preference"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/add-notification-preferences [patch]
func (h *Handler) AddNotificationPreference(c *gin.Context) {
	var input *dtos.NotificationPreference
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.AddNotificationPreference(c, input, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Push Notification
// @Description  Toggle Push Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-push-notification [patch]
func (h *Handler) TogglePushNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.TogglePushNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Likes Notification
// @Description  Toggle Likes Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-likes-notification [patch]
func (h *Handler) ToggleLikesNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.TogglePushNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Comments Notification
// @Description  Toggle Comments Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-comment-notification [patch]
func (h *Handler) ToggleCommentsNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.ToggleCommentsNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Tags and Mentions Notification
// @Description Toggle Tags and Mentions Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-tags-and-mention-notification [patch]
func (h *Handler) ToggleTagsAndMentionNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.ToggleTagsAndMentionsNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Repost Notification
// @Description Toggle Repost Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-repost-notification [patch]
func (h *Handler) ToggleRepostNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.ToggleRepostNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Direct Message Notification
// @Description Toggle Direct Message Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-direct-message-notification [patch]
func (h *Handler) ToggleDirectMessageNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.ToggleDirectMessageNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle Live Notification
// @Description Toggle Live Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-live-notification [patch]
func (h *Handler) ToggleLiveNotification(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.ToggleLiveNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Toggle New Follower Notification
// @Description Toggle  New Follower Notification
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/toggle-new-follower-notification [patch]
func (h *Handler) ToggleNewFollower(c *gin.Context) {
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.ToggleNewFollowerNotification(c, &user)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Set Content Settings
// @Description Set Content Settings
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/content-settings [patch]
func (h *Handler) SetContentSettings(c *gin.Context) {

	var input dtos.ContentSettingsDto
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.SetContentSettings(c, &user, &input)
	c.JSON(result.Code, result)
}

// @Tags Settings
// @Summary Set Banned Words
// @Description Set Banned Words
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /settings/banned-words [patch]
func (h *Handler) SetBannedWords(c *gin.Context) {
	var input dtos.SetBannedWordsDto
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to core
	result := h.core.SetBannedWords(c, &input, &user)
	c.JSON(result.Code, result)
}
