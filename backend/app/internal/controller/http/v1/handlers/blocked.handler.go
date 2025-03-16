package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"backend.app/common/constants"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// @Tags Block
// @Summary Get Blocked Users
// @Description Gets Blocked Users
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.APIPagingDto   true  "data to query for all "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /block [get]
func (h *Handler) GetAllBlockedUsers(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.GetAllBlockedUsers(c, &user, &query)
	c.JSON(result.Code, result)
}

// @Tags Block
// @Summary Get Blocked User
// @Description Gets Blocked User
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "The Id of the blocked user"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /block/{id} [get]
func (h *Handler) GetBlockedUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.GetBlockedUser(c, id, &user)
	c.JSON(result.Code, result)
}
