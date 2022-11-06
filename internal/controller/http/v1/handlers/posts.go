package handlers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postsRoute struct {
	uc     usecase.Post
	logger logger.Interface
	mapper mappers.PostMapper
}

func NewPostsRoute(handler *gin.RouterGroup, i usecase.Post, l logger.Interface, m mappers.PostMapper) {
	c := &postsRoute{i, l, m}

	handler.POST("/posts", c.create)
	handler.DELETE("/posts/:postID", c.delete)
	handler.GET("/posts", c.get)

}

type createPostResponse struct {
	post dtos.Post `json:"post"`
}

type createPostRequest struct {
	Message string `json:"message"  binding:"required"  example:"What is happening?"`
	Media   []struct {
		URL           string `json:"url"  binding:"required"`
		AlternateText string `json:"alternate_text"`
	} `json:"media"`
}

type getPostsRequest struct {
	Page  int
	Limit int
}

type getPostsResponse struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Posts []dtos.Post `json:"posts"`
}

// @Summary     Create Post
// @Description Create a post
// @ID          create-post
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request body     createPostRequest true "create a new post"
// @Success     201     {object} createPostResponse
// @Router      /posts [post]
func (r *postsRoute) create(c *gin.Context) {
	var request createPostRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()
	post := &models.Post{
		Message: &request.Message,
	}

	err := r.uc.Create(ctx, post)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		HTTPError(c, http.StatusInternalServerError, "unable to create a post")

		return
	}

	c.JSON(http.StatusCreated, createPostResponse{})
}

// @Summary     Get Posts
// @Description Get user's posts
// @ID          get-posts
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       request query    getPostsRequest true "create a new post"
// @Success     200     {object} getPostsResponse
// @Router      /posts [get]
func (r *postsRoute) get(c *gin.Context) {
	var request getPostsRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	//ctx := c.Request.Context()
	//
	//
	//err := r.uc.Create(ctx, post)
	//if err != nil {
	//	r.logger.Error(err, "http - v1 - registerUser")
	//	HTTPError(c, http.StatusInternalServerError, "unable to create a post")
	//
	//	return
	//}

	c.JSON(http.StatusCreated, getPostsResponse{})
}

type deletePostResponse struct {
	Message string `json:"message"`
}

// @Summary     Delete Post
// @Description Delete a post
// @ID          delete-post
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       postID path     string true "post post identifier"
// @Success     200           {object} deletePostResponse
// @Failure     400           {object} ErrorResponse
// @Failure     401           {object} ErrorResponse
// @Failure     404           {object} ErrorResponse
// @Failure     500           {object} ErrorResponse
// @Router      /posts/{postID} [delete]
func (r *postsRoute) delete(c *gin.Context) {
	cID := c.Param("postID")

	post, err := r.uc.Get(c.Request.Context(), cID)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		HTTPError(c, http.StatusNotFound, "unable to delete post")

		return
	}

	// TODO: VALIDATE USER HAS RIGHT TO DELETE POSTS

	err = r.uc.Delete(c.Request.Context(), post)
	if err != nil {
		r.logger.Error(err, "http - v1 - posts - create")
		HTTPError(c, http.StatusInternalServerError, "unable to delete post")

		return
	}

	c.JSON(http.StatusOK, deletePostResponse{
		Message: "post has been successfully deleted",
	})
}
