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

type timelineRoute struct {
	usersService userservice.User
	postsService posts.Post
	logger       logger.Interface
	mapper       mappers.PostMapper
}

func NewTimelineRoute(handler *gin.RouterGroup, u userservice.User, p posts.Post, l logger.Interface, m mappers.PostMapper) {
	c := &timelineRoute{u, p, l, m}

	handler.GET("/timeline", c.get)
}

type getTimelineRequest struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type getTimelineResponse struct {
	Status string            `json:"status"`
	Data   dtos.PagedContent `json:"data"`
}

// @Summary     Get Timeline
// @Description Get session user's timeline
// @ID          get-timeline
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request query    getTimelineRequest true "get user timeline"
// @Success     200     {object} getTimelineResponse
// @Security    Bearer
// @Router      /timeline [get]
func (r *timelineRoute) get(c *gin.Context) {
	var request getTimelineRequest

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
	_, err = r.postsService.GetAll(ctx, sessionUser.ID, request.Page, request.Limit)
	if err != nil {
		r.logger.Error(err, "http - v1 -  timeline - get")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to get timeline")

		return
	}

	c.JSON(http.StatusOK, getTimelineResponse{
		Status: "success",
	})
}
