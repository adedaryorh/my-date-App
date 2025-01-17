package relationships

import (
	"net/http"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/internal/controller/response"
	"backend.app/internal/dtos"
	"backend.app/internal/services/relationships"
	"backend.app/pkg/logger"
	"github.com/gin-gonic/gin"
)

type lovedOnesRoute struct {
	lovedOneService relationships.LovedOne
	logger          logger.Interface
}

func NewLovedOnesRoute(handler *gin.RouterGroup, r relationships.LovedOne, l logger.Interface) {
	c := &lovedOnesRoute{r, l}

	handler.POST("/lovedones/discover", c.discover)
	handler.POST("/lovedones/:userID", c.becomeLovedOne)
	handler.DELETE("/lovedones/:userID", c.removeLovedOne)
	handler.GET("/lovedones", c.getLovedOnes)
	//handler.POST("/followers/report/:userID", c.report)
}

type discoverLovedOnesRequest struct {
	Contacts []string `json:"contacts"`
}

type discoverLovedOnesResponse struct {
	Status   string          `json:"status"`
	Contacts []dtos.UserInfo `json:"contacts"`
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
	businessID := c.Param("userID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - loved ones - add")
		handlers.HTTPError(c, http.StatusUnauthorized, "unable to get session user")

		return
	}

	err = r.lovedOneService.AddLovedOne(c.Request.Context(), sessionUser.ID, businessID)
	if err != nil {
		r.logger.Error(err, "http - v1 - loved ones - add")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			errorMessage := "unable to add loved one"
			if len(serviceErr.Message) > 0 {
				errorMessage = serviceErr.Message
			}

			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, errorMessage, serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to add loved one")
		}

		return
	}

	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}

// @Summary     Discover loved ones
// @Description Discover loved ones
// @ID          discover-loved-ones
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       request body   discoverLovedOnesRequest   true "contacts"
// @Success     200     {object} discoverLovedOnesResponse
// @Security Bearer
// @Router      /lovedones/discover [post]
func (r *lovedOnesRoute) discover(c *gin.Context) {
	var request discoverLovedOnesRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - loved ones - discover")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	lo, err := r.lovedOneService.DiscoverLovedOnes(c.Request.Context(), request.Contacts)
	if err != nil {
		r.logger.Error(err, "http - v1 - loved ones - discover")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			errorMessage := "error discovering loved one"
			if len(serviceErr.Message) > 0 {
				errorMessage = serviceErr.Message
			}

			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, errorMessage, serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "error discovering loved one")
		}

		return
	}

	c.JSON(http.StatusOK, discoverLovedOnesResponse{
		Status:   "success",
		Contacts: lo,
	})
}

// @Summary     Get loved ones
// @Description Get user's loved ones
// @ID          get-loved-ones
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       request query    getFollowersRequest true "get user's loved ones"
// @Success     200     {object} getCustomersResponse
// @Security    Bearer
// @Router      /lovedones [get]
func (r *lovedOnesRoute) getLovedOnes(c *gin.Context) {
	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - loved ones - get")
		handlers.HTTPError(c, http.StatusUnauthorized, "unable to get session user")

		return
	}

	data, err := r.lovedOneService.GetLovedOnes(c.Request.Context(), sessionUser.ID)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - get")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to get loved ones", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to get loved ones")
		}

		return
	}

	c.JSON(http.StatusOK, getCustomersResponse{
		Status: "success",
		Data:   *data,
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
	businessID := c.Param("userID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - loved ones - remove")
		handlers.HTTPError(c, http.StatusUnauthorized, "unable to get session user")

		return
	}

	err = r.lovedOneService.RemoveLovedOne(c.Request.Context(), sessionUser.ID, businessID)
	if err != nil {
		r.logger.Error(err, "http - v1 - loved ones - remove")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to remove loved one", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to remove loved one")
		}

		return
	}

	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}
