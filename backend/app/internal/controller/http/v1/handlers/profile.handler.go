package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/common/messages"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
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
// @Router /profile [put]
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	var input *dtos.UpdateUserProfile
	// bind input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  constants.HttpStatusBadRequest,
			"error":   err,
			"message": messages.ErrInvalidInput.Error(),
		})
		return
	}
	if inputErrors := helpers.ValidateInput(input); inputErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  constants.HttpStatusBadRequest,
			"error":   inputErrors,
			"message": messages.ErrInvalidInput.Error(),
		})
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
// @Router /profile/add-profile-image [put]
func (h *Handler) UploadUserProfilePicture(c *gin.Context) {
	var input *dtos.UploadImage
	// bind input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  constants.HttpStatusBadRequest,
			"error":   err,
			"message": messages.ErrInvalidInput.Error(),
		})
		return
	}
	if inputErrors := helpers.ValidateInput(input); inputErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  constants.HttpStatusBadRequest,
			"error":   inputErrors,
			"message": messages.ErrInvalidInput.Error(),
		})
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.UploadUserProfileImage(c, &user, input)
	c.JSON(result.Code, result)
}
