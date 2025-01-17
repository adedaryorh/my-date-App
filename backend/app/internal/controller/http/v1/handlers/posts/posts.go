package posts

import (
	"net/http"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/internal/controller/response"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/services/posts"
	userservice "backend.app/internal/services/users"
	"backend.app/internal/validators"
	"backend.app/pkg/logger"
	"github.com/gin-gonic/gin"
)

type postsRoute struct {
	usersService   userservice.User
	postsService   posts.Post
	logger         logger.Interface
	mapper         mappers.PostMapper
	postsValidator validators.ValidatePost
}

func NewPostsRoute(handler *gin.RouterGroup, u userservice.User, p posts.Post, l logger.Interface, m mappers.PostMapper) {
	c := &postsRoute{u, p, l, m, validators.NewPostValidator()}

	handler.POST("/posts", c.create)
	handler.DELETE("/posts/:postID", c.delete)
	handler.GET("/posts", c.get)
	handler.GET("/posts/:userID", c.getUserPosts)
	handler.POST("/posts/report/:postID", c.report)
}

type createPostResponse struct {
	Status string    `json:"status"`
	Post   dtos.Post `json:"post"`
}

type createPostRequest struct {
	Message string   `json:"message" example:"What is happening?"`
	Media   []string `form:"media" json:"media" `
}

type getPostsRequest struct {
	Page  *int `form:"page" json:"page"`
	Limit *int `form:"limit" json:"limit"`
}

type getPostsResponse struct {
	Status string          `json:"status"`
	Data   dtos.PagedPosts `json:"data"`
}

type reportPostResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// @Summary     Create Post
// @Description Create a post
// @ID          create-post
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request body createPostRequest true "create a new post"
// @Success     201     {object} createPostResponse
// @Security Bearer
// @Router      /posts [post]
func (r *postsRoute) create(c *gin.Context) {
	var request createPostRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to create a post")
	}

	ctx := c.Request.Context()

	post := dtos.NewPost{
		Message: &request.Message,
		Author:  *sessionUser,
		Media:   request.Media,
	}

	err = r.postsValidator.Validate(post)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "post validation failed", err)

		return
	}

	dto, err := r.postsService.Create(ctx, post)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to create a post", err)

		return
	}

	c.JSON(http.StatusCreated, createPostResponse{
		Status: "success",
		Post:   *dto,
	})
}

// @Summary     Get Posts
// @Description Get session user's posts
// @ID          get-posts
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request query    getPostsRequest true "get user posts"
// @Success     200     {object} getPostsResponse
// @Security    Bearer
// @Router      /posts [get]
func (r *postsRoute) get(c *gin.Context) {
	var request getPostsRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - posts - get")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to create a post")
	}

	ctx := c.Request.Context()
	p, err := r.postsService.GetAll(ctx, sessionUser.ID, request.Page, request.Limit)
	if err != nil {
		r.logger.Error(err, "http - v1 -  posts - get")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to get posts")

		return
	}

	c.JSON(http.StatusOK, getPostsResponse{
		Status: "success",
		Data:   *p,
	})
}

type deletePostResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// @Summary     Delete Post
// @Description Delete a post
// @ID          delete-post
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       postID path     string true "post identifier"
// @Success     200           {object} deletePostResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /posts/{postID} [delete]
func (r *postsRoute) delete(c *gin.Context) {
	postID := c.Param("postID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to create a post")
	}

	err = r.postsService.Delete(c.Request.Context(), sessionUser.ID, postID)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to delete post", err)

		return
	}

	c.JSON(http.StatusOK, deletePostResponse{
		Status:  "success",
		Message: "Post has been successfully deleted",
	})
}

// @Summary     Report post
// @Description Report a post
// @ID          report-post
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       postID path     string true "post identifier"
// @Success     200           {object} reportPostResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /posts/report/{postID} [post]
func (r *postsRoute) report(c *gin.Context) {
	postID := c.Param("postID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - report")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to create a post")

		return
	}

	err = r.postsService.Report(c.Request.Context(), sessionUser.ID, postID)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - report")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to report post", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to report post")
		}

		return
	}

	c.JSON(http.StatusOK, reportPostResponse{
		Status:  "success",
		Message: "Thanks for reporting. We will get back to you.",
	})
}

// @Summary     Get User Posts by ID
// @Description Get user's posts by user ID
// @ID          get-user-posts
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request query   getPostsRequest true "get user posts by ID"
// @Param       userID path     string true "user identifier"
// @Success     200     {object} getPostsResponse
// @Security    Bearer
// @Router      /posts/{userID} [get]
func (r *postsRoute) getUserPosts(c *gin.Context) {
	userID := c.Param("userID")

	var request getPostsRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - posts - get")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()

	user, err := r.usersService.ValidateUser(ctx, userID)
	if err != nil {
		r.logger.Error(err, "http - v1 -  posts - get")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to get user")

		return
	}

	p, err := r.postsService.GetAll(ctx, user.ID, request.Page, request.Limit)
	if err != nil {
		r.logger.Error(err, "http - v1 -  posts - get")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to get user posts")

		return
	}

	c.JSON(http.StatusOK, getPostsResponse{
		Status: "success",
		Data:   *p,
	})
}
