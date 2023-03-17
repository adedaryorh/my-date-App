package celebrations

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/controller/response"
	"celebut-api/internal/dtos"
	"celebut-api/internal/services/celebrations"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetCelebrationsType string

const (
	Recent    GetCelebrationsType = "recent"
	AroundYou GetCelebrationsType = "around-you"
	LovedOnes GetCelebrationsType = "loved-ones"
)

type celebrationsRoute struct {
	usersService        userservice.User
	celebrationsService celebrations.Celebration
	logger              logger.Interface
}

func NewCelebrationsRoute(handler *gin.RouterGroup, u userservice.User, p celebrations.Celebration, l logger.Interface) {
	c := &celebrationsRoute{u, p, l}

	handler.GET("/celebrations", c.get)
	handler.POST("/celebrations/report/:celebrationID", c.report)
}

type createCelebrationResponse struct {
	Status      string           `json:"status"`
	Celebration dtos.Celebration `json:"data"`
}

type createCelebrationRequest struct {
	Caption string   `json:"caption" example:"celebration caption"`
	Media   []string `json:"media"`
}

type getCelebrationsRequest struct {
	Page  *int                `form:"page" json:"page"`
	Type  GetCelebrationsType `form:"type" json:"type"  binding:"required"`
	Limit *int                `form:"limit" json:"limit"`
}

type getCelebrationsResponse struct {
	Status string                 `json:"status"`
	Data   dtos.PagedCelebrations `json:"data"`
}

type reportCelebrationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type deleteCelebrationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// @Summary     Create Celebration
// @Description Create a celebration
// @ID          create-celebration
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       request body createCelebrationRequest true "create a new celebration"
// @Success     201     {object} createCelebrationResponse
// @Security Bearer
// @Router      /celebrations [post]
func (r *celebrationsRoute) create(c *gin.Context) {
	var request createCelebrationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "invalid request body", err)

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "unable to create a celebration", err)

		return
	}

	ctx := c.Request.Context()

	celebration := dtos.NewCelebration{
		Message: &request.Caption,
		Author:  *sessionUser,
		Media:   request.Media,
	}

	dto, err := r.celebrationsService.Create(ctx, celebration)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to create a celebration", err)

		return
	}

	c.JSON(http.StatusCreated, createCelebrationResponse{
		Status:      "success",
		Celebration: *dto,
	})
}

// @Summary     Get Celebrations
// @Description Get celebrations based on request type
// @ID          get-celebrations
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       request query    getCelebrationsRequest true "get celebrations params"
// @Param       request query    getCelebrationsRequest true "get celebrations params"
// @Success     200     {object} getCelebrationsResponse
// @Security    Bearer
// @Router      /celebrations [get]
func (r *celebrationsRoute) get(c *gin.Context) {
	var request getCelebrationsRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - celebrations - get")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - get")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "unable to get celebrations", err)
	}

	ctx := c.Request.Context()

	var pc *dtos.PagedCelebrations

	if request.Type == Recent {
		pc, err = r.celebrationsService.Recent(ctx, sessionUser.ID, request.Page, request.Limit)
	} else if request.Type == AroundYou {
		pc, err = r.celebrationsService.AroundYou(ctx, sessionUser.ID, request.Page, request.Limit)
	} else if request.Type == LovedOnes {
		pc, err = r.celebrationsService.LovedOnes(ctx, sessionUser.ID, request.Page, request.Limit)
	} else {
		err = errors.New("invalid celebrations type")
	}

	if err != nil {
		r.logger.Error(err, "http - v1 -  celebrations - recent")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to get celebrations", err)

		return
	}

	c.JSON(http.StatusOK, getCelebrationsResponse{
		Status: "success",
		Data:   *pc,
	})
}

// @Summary     Delete Celebration
// @Description Delete a celebration
// @ID          delete-celebration
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       celebrationID path     string true "celebration identifier"
// @Success     200           {object} deleteCelebrationResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /celebrations/{celebrationID} [delete]
func (r *celebrationsRoute) delete(c *gin.Context) {
	celebrationID := c.Param("celebrationID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "unable to delete a celebration", err)

		return
	}

	err = r.celebrationsService.Delete(c.Request.Context(), sessionUser.ID, celebrationID)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		handlers.HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to delete a celebration", err)

		return
	}

	c.JSON(http.StatusOK, deleteCelebrationResponse{
		Status:  "success",
		Message: "Celebration has been successfully deleted",
	})
}

// @Summary     Report celebration
// @Description Report a celebration
// @ID          report-celebration
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       celebrationID path     string true "celebration identifier"
// @Success     200           {object} reportCelebrationResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /celebrations/report/{celebrationID} [post]
func (r *celebrationsRoute) report(c *gin.Context) {
	celebrationID := c.Param("celebrationID")

	sessionUser, err := handlers.GetSessionUser(c)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - report")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "unable to report celebrations", err)

		return
	}

	err = r.celebrationsService.Report(c.Request.Context(), sessionUser.ID, celebrationID)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - report")

		serviceErr, ok := err.(*response.ServiceErrorResponse)
		if ok {
			handlers.HTTPErrorWithInformation(c, serviceErr.StatusCode, "unable to report celebrations", serviceErr.Err)
		} else {
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to report celebrations")
		}

		return
	}

	c.JSON(http.StatusOK, reportCelebrationResponse{
		Status:  "success",
		Message: "Thanks for reporting. We will get back to you.",
	})
}
