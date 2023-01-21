package posts

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/services/posts"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type feedRoute struct {
	usersService userservice.User
	postsService posts.Post
	logger       logger.Interface
	mapper       mappers.PostMapper
}

func NewFeedRoute(handler *gin.RouterGroup, u userservice.User, p posts.Post, l logger.Interface, m mappers.PostMapper) {
	c := &feedRoute{u, p, l, m}

	handler.GET("/posts/feed", c.get)
}

type getFeedRequest struct {
	Page  *int `form:"page" json:"page"`
	Limit *int `form:"limit" json:"limit"`
}

type getFeedResponse struct {
	Status string          `json:"status"`
	Data   dtos.PagedPosts `json:"data"`
}

// @Summary     Get Posts Feed
// @Description Get session user's feed
// @ID          get-posts-feed
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request query    getFeedRequest true "get user timeline"
// @Success     200     {object} getFeedResponse
// @Security    Bearer
// @Router      /posts/feed [get]
func (r *feedRoute) get(c *gin.Context) {
	var request getFeedRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - timeline - get")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - timeline - create")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to get timeline")

		return
	}

	ctx := c.Request.Context()
	feed, err := r.postsService.GetFeed(ctx, sessionUser.ID, request.Page, request.Limit)
	if err != nil {
		r.logger.Error(err, "http - v1 -  timeline - get")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to get timeline")

		return
	}

	c.JSON(http.StatusOK, getFeedResponse{
		Status: "success",
		Data:   *feed,
	})
}
