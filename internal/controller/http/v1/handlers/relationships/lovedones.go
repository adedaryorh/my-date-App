package relationships

import (
	"celebut-api/internal/dtos"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type lovedOnesRoute struct {
	usersService userservice.User
	logger       logger.Interface
}

func NewLovedOnesRoute(handler *gin.RouterGroup, u userservice.User, l logger.Interface) {
	c := &lovedOnesRoute{u, l}

	handler.POST("/lovedones/:userID", c.becomeLovedOne)
	handler.DELETE("/lovedones/:userID", c.removeLovedOne)
	handler.GET("/lovedones", c.getLovedOnes)
	//handler.POST("/followers/report/:userID", c.report)
}

type getLovedOnesRequest struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type getLovedOnesResponse struct {
	Status string                  `json:"status"`
	Data   dtos.PagedRelationships `json:"data"`
}

type lovedOnesResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// @Summary     Become a loved one
// @Description Become a loved one
// @ID          become-a-loved-one
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       userID path     string true "user identifier"
// @Success     200     {object} followerResponse
// @Security Bearer
// @Router      /lovedones/{userID} [post]
func (r *lovedOnesRoute) becomeLovedOne(c *gin.Context) {
	// TODO: Do followers logic
	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}

// @Summary     Get loved ones
// @Description Get user's loved ones
// @ID          get-loved-ones
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       request query    getFollowersRequest true "get user's loved ones"
// @Success     200     {object} getFollowersResponse
// @Security    Bearer
// @Router      /lovedones [get]
func (r *lovedOnesRoute) getLovedOnes(c *gin.Context) {

	// TODO: Do followers logic
	c.JSON(http.StatusOK, getFollowersResponse{
		Status: "success",
	})
}

// @Summary     Remove loved one
// @Description Remove user's loved one
// @ID          remove-loved-one
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       userID path     string true "user identifier"
// @Success     200           {object} followerResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /lovedones/{userID} [delete]
func (r *lovedOnesRoute) removeLovedOne(c *gin.Context) {
	// TODO: Do unfollowers logic

	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}
