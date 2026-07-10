package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// @Tags Profile
// @Summary Update user information
// @Description Updates user information
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.UpdateUserProfile  true  "data to update user personal info "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /profile [patch]
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	var input *dtos.UpdateUserProfile
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
	result := h.core.UpdateUserProfile(c, &user, input)
	c.JSON(result.Code, result)
}

// @Tags Profile
// @Summary Add User Profile Picture
// @Schemes
// @Description Add User Profile Picture
// @Param   request   body     dtos.UploadImage   true  "data to add user profile picture"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "desc"
// @Router /profile/add-profile-image [patch]
func (h *Handler) UploadUserProfilePicture(c *gin.Context) {
	var input *dtos.UploadImage
	// bind input
	if err := c.Bind(&input); err != nil {
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
	result := h.core.UploadUserProfileImage(c, &user, input)
	c.JSON(result.Code, result)
}

// @Tags Profile
// @Summary Block User
// @Description Blocks User
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "The Id of the User to be blocked"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /profile/{id}/block [patch]
func (h *Handler) BlockUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.BlockUser(c, id, &user)
	c.JSON(result.Code, result)
}

// @Tags Profile
// @Summary Un Block User
// @Description Un Blocks User
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "The Id of the User to be un blocked"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /profile/{id}/unblock [patch]
func (h *Handler) UnBlockUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.UnBlockUser(c, id, &user)
	c.JSON(result.Code, result)
}

// @Tags Profile
// @Summary Follow User
// @Description Follows User
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "The Id of the User to be followed"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /profile/{id}/follow [patch]
func (h *Handler) FollowUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.FollowUser(c, id, &user)
	c.JSON(result.Code, result)
}

// @Tags Profile
// @Summary UnFollow User
// @Description UnFollows User
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "The Id of the User to be unollowed"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /profile/{id}/unfollow [patch]
func (h *Handler) UnFollowUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.UnFollowUser(c, id, &user)
	c.JSON(result.Code, result)
}
