package relationships

import (
	"celebut-api/internal/dtos"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type followersRoute struct {
	usersService userservice.User
	logger       logger.Interface
}

func NewFollowersRoute(handler *gin.RouterGroup, u userservice.User, l logger.Interface) {
	c := &followersRoute{u, l}

	handler.POST("/follow/:userID", c.follow)
	handler.DELETE("/unfollow/:userID", c.unfollow)
	handler.GET("/followers", c.get)
	//handler.POST("/followers/report/:userID", c.report)
}

type getFollowersRequest struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type getFollowersResponse struct {
	Status string                  `json:"status"`
	Data   dtos.PagedRelationships `json:"data"`
}

type followerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// @Summary     Follow business
// @Description Follow a business
// @ID          follow-business
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       userID path     string true "user identifier"
// @Success     200     {object} followerResponse
// @Security Bearer
// @Router      /followers/{userID} [post]
func (r *followersRoute) follow(c *gin.Context) {
	// TODO: Do followers logic
	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}

// @Summary     Get followers
// @Description Get businesses' followers
// @ID          get-followers
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       request query    getFollowersRequest true "get businesses' followers"
// @Success     200     {object} getFollowersResponse
// @Security    Bearer
// @Router      /followers [get]
func (r *followersRoute) get(c *gin.Context) {

	// TODO: Do followers logic
	c.JSON(http.StatusOK, getFollowersResponse{
		Status: "success",
	})
}

// @Summary     Unfollow business
// @Description Unfollow a business
// @ID          unfollow-business
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
// @Router      /unfollow/{userID} [delete]
func (r *followersRoute) unfollow(c *gin.Context) {
	// TODO: Do unfollowers logic

	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}
