package handlers

import (
	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
	"github.com/gin-gonic/gin"
)

// @Tags Celebration
// @Summary Create Celebration
// @Schemes
// @Description Create Celebration
// @Param   request   body     dtos.UploadImage   true  "data to add user profile picture"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.CreateCelebrationDto "desc"
// @Router /celebrations [post]
func (h *Handler) CreateCelebration(c *gin.Context) {
	var input *dtos.CreateCelebrationDto
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
	result := h.core.CreateCelebrationDto(c, &user, input)
	c.JSON(result.Code, result)
}

// @Tags Celebration
// @Summary Get Celebrations
// @Description Get Celebrations
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.APIPagingDto   true  "data to query for all "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /celebrations [get]
func (h *Handler) GetCelebrations(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.GetAllCelebrations(c, &user, &query)
	c.JSON(result.Code, result)
}
