package posts

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/services/posts"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type reactionsRoute struct {
	usersService userservice.User
	postsService posts.Post
	logger       logger.Interface
	mapper       mappers.PostMapper
}

func NewReactionsRoute(handler *gin.RouterGroup, u userservice.User, p posts.Post, l logger.Interface, m mappers.PostMapper) {
	c := &reactionsRoute{u, p, l, m}

	handler.POST("/posts/reactions/:postID", c.addReaction)
	handler.DELETE("/posts/reactions/:postID", c.deleteReaction)
}

type addReactionResponse struct {
	Status string    `json:"status"`
	Post   dtos.Post `json:"post"`
}

type addReactionRequest struct {
	Reaction string                 `json:"reaction" example:"emoji"`
	Media    []multipart.FileHeader `form:"media" json:"media" swaggerignore:"true"`
}

type deleteReactionResponse struct {
	Status string `json:"status"`
}

// @Summary     Add reaction to a post
// @Description Add reaction to a post
// @ID          add-post-reaction
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       postID path     string true "post identifier"
// @Param       request body addReactionRequest true "add a new reaction"
// @Success     200     {object} addReactionResponse
// @Security Bearer
// @Router      /posts/reactions/{postID} [post]
func (r *reactionsRoute) addReaction(c *gin.Context) {
	postID := c.Param("postID")
	var request addReactionRequest

	if err := c.ShouldBind(&request); err != nil {
		r.logger.Error(err, "http - v1 - reactions - add")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "invalid request body", err)

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - reactions - add")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "unable to add a reaction", err)

		return
	}

	ctx := c.Request.Context()
	reaction := dtos.NewReaction{
		Reaction: request.Reaction,
		User:     *sessionUser,
	}

	dto, err := r.postsService.AddReaction(ctx, postID, reaction)
	if err != nil {
		r.logger.Error(err, "http - v1 - reactions - add")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to add a reaction", err)

		return
	}

	c.JSON(http.StatusCreated, addReactionResponse{
		Status: "success",
		Post:   *dto,
	})
}

// @Summary     Remove reaction
// @Description Remove reaction from a post
// @ID          delete-post-reaction
// @Tags        Posts
// @Accept      json
// @Produce     json
// @Param       postID path     string true "post identifier"
// @Success     200     {object} addReactionResponse
// @Security Bearer
// @Router      /posts/reactions/{postID} [delete]
func (r *reactionsRoute) deleteReaction(c *gin.Context) {
	postID := c.Param("postID")
	var request addReactionRequest

	if err := c.ShouldBind(&request); err != nil {
		r.logger.Error(err, "http - v1 - reactions - delete")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "invalid request body", err)

		return
	}

	user, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - reactions - delete")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "unable to add reaction", err)

		return
	}

	ctx := c.Request.Context()
	err = r.postsService.DeleteReaction(ctx, user.ID, postID)
	if err != nil {
		r.logger.Error(err, "http - v1 - reactions - delete")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to remove reaction", err)

		return
	}

	c.JSON(http.StatusOK, deleteReactionResponse{
		Status: "success",
	})
}
