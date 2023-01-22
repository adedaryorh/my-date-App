package relationships

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/controller/response"
	"celebut-api/internal/dtos"
	"celebut-api/internal/services/relationships"

	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type customersRoute struct {
	customersService relationships.Customer
	logger           logger.Interface
}

func NewCustomersRoute(handler *gin.RouterGroup, f relationships.Customer, l logger.Interface) {
	c := &customersRoute{f, l}

	handler.POST("/follow/:businessID", c.follow)
	handler.DELETE("/unfollow/:businessID", c.unfollow)
	handler.GET("/followers", c.get)
	//handler.POST("/followers/report/:userID", c.report)
}

type getFollowersRequest struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

type getCustomersResponse struct {
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
// @Param       businessID path     string true "business user identifier"
// @Success     200     {object} followerResponse
// @Security Bearer
// @Router      /follow/{businessID} [post]
func (r *customersRoute) follow(c *gin.Context) {
	businessID := c.Param("businessID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - follow")
		handlers.HTTPError(c, http.StatusUnauthorized, "unable to get session user")

		return
	}

	err = r.customersService.FollowBusiness(c.Request.Context(), sessionUser.ID, businessID)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - follow")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			errorMessage := "unable to follow user"
			if len(serviceErr.Message) > 0 {
				errorMessage = serviceErr.Message
			}

			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, errorMessage, serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to follow user")
		}

		return
	}

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
// @Success     200     {object} getCustomersResponse
// @Security    Bearer
// @Router      /followers [get]
func (r *customersRoute) get(c *gin.Context) {
	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - get")
		handlers.HTTPError(c, http.StatusUnauthorized, "unable to get session user")

		return
	}

	data, err := r.customersService.GetFollowers(c.Request.Context(), sessionUser.ID)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - get")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to get followers", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to get followers")
		}

		return
	}

	c.JSON(http.StatusOK, getCustomersResponse{
		Status: "success",
		Data:   *data,
	})
}

// @Summary     Unfollow business
// @Description Unfollow a business
// @ID          unfollow-business
// @Tags        User Relationships
// @Accept      json
// @Produce     json
// @Param       businessID path     string true "business user identifier"
// @Success     200           {object} followerResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /unfollow/{businessID} [delete]
func (r *customersRoute) unfollow(c *gin.Context) {
	businessID := c.Param("businessID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - unfollow")
		handlers.HTTPError(c, http.StatusUnauthorized, "unable to get session user")

		return
	}

	err = r.customersService.UnfollowBusiness(c.Request.Context(), sessionUser.ID, businessID)
	if err != nil {
		r.logger.Error(err, "http - v1 - followers - unfollow")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to follow user", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to follow user")
		}

		return
	}

	c.JSON(http.StatusOK, followerResponse{
		Status:  "success",
		Message: "success",
	})
}
