package posts

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/controller/response"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/services/posts"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type commentsRoute struct {
	usersService userservice.User
	postsService posts.Post
	logger       logger.Interface
	mapper       mappers.PostMapper
}

func NewCommentsRoute(handler *gin.RouterGroup, u userservice.User, p posts.Post, l logger.Interface, m mappers.PostMapper) {
	c := &commentsRoute{u, p, l, m}

	handler.POST("/comments/:postID", c.comment)
	handler.GET("/comments/:postID", c.get)
}

type createCommentResponse struct {
	Status string    `json:"status"`
	Post   dtos.Post `json:"post"`
}

type createCommentRequest struct {
	Message string   `json:"message" example:"What is happening?"`
	Media   []string `json:"media"`
}

type getCommentsRequest struct {
	Page  *int `form:"page" json:"page"`
	Limit *int `form:"limit" json:"limit"`
}

type getCommentsResponse struct {
	Status string         `json:"status"`
	Data   getCommentData `json:"data"`
}

type getCommentData struct {
	Post     dtos.Post       `json:"post"`
	Comments dtos.PagedPosts `json:"comments"`
}

// @Summary     Create post comment
// @Description Add comment to a post
// @ID          create-comment
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       postID path     string true "post identifier"
// @Param       request body createCommentRequest true "create a new comment"
// @Success     201     {object} createCommentResponse
// @Security Bearer
// @Router      /comments/{postID} [post]
func (r *commentsRoute) comment(c *gin.Context) {
	postID := c.Param("postID")
	var request createCommentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - comments - create")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - comments - create")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to add a comment")

		return
	}

	ctx := c.Request.Context()

	post, err := r.postsService.ValidatePost(ctx, postID)
	if err != nil {
		r.logger.Error(err, "http - v1 - comments - create")
		handlers.HTTPError(c, http.StatusBadRequest, "post not found")

		return
	}

	comment := dtos.NewPost{
		Message: &request.Message,
		Author:  *sessionUser,
		Media:   request.Media,
	}

	dto, err := r.postsService.Comment(ctx, post.ID, comment)
	if err != nil {
		r.logger.Error(err, "http - v1 - comments - create")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to create a post")

		return
	}

	c.JSON(http.StatusCreated, createCommentResponse{
		Status: "success",
		Post:   *dto,
	})
}

// @Summary     Get Post Comments
// @Description Get Post Comments
// @ID          get-post-comments
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request query    getCommentsRequest true "get post comments"
// @Param       postID path     string true "post identifier"
// @Success     200     {object} getCommentsResponse
// @Security    Bearer
// @Router      /comments/{postID} [get]
func (r *commentsRoute) get(c *gin.Context) {
	var request getCommentsRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - posts - get")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()
	post, comments, err := r.postsService.GetComments(ctx, c.Param("postID"), request.Page, request.Limit)
	if err != nil {
		r.logger.Error(err, "http - v1 -  posts - get")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to get comments", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to get posts")
		}

		return
	}

	c.JSON(http.StatusOK, getCommentsResponse{
		Status: "success",
		Data:   getCommentData{Post: *post, Comments: *comments},
	})
}
