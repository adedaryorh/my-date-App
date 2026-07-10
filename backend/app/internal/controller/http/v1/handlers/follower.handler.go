package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"backend.app/common/constants"
	"backend.app/internal/models"
	"backend.app/pkg/response"
)

// @Tags Followers
// @Summary Get All Followers
// @Description Get All Followers
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.APIPagingDto   true  "data to query for all "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /followers [get]
func (h *Handler) GetAllFollowers(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.GetAllFollowers(c, &user, &query)
	c.JSON(result.Code, result)
}

// @Tags Followers
// @Summary Get Friends
// @Description Get Friends
// @Accept  json
// @Produce  json
// @Param   request   body     dtos.APIPagingDto   true  "data to query for all "
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /followers/friends [get]
func (h *Handler) GetFriends(c *gin.Context) {
	query := getPagingInfo(c)
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.GetFriends(c, &user, &query)
	c.JSON(result.Code, result)
}

// @Tags Followers
// @Summary Get Single Follower
// @Description Gets Single Follower
// @Accept  json
// @Produce  json
// @Param   id   path     string   true  "The Id of the blocked user"
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /followers/{id} [get]
func (h *Handler) GetSingleFollower(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	user := c.MustGet("authUser").(models.User) // auth user
	// send to controller
	result := h.core.GetSingleFollower(c, id, &user)
	c.JSON(result.Code, result)
}
