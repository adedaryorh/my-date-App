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

type celebrationsRoute struct {
	uc     usecase.Celebration
	logger logger.Interface
	mapper mappers.CelebrationMapper
}

func NewCelebrationsRoute(handler *gin.RouterGroup, i usecase.Celebration, l logger.Interface, m mappers.CelebrationMapper) {
	c := &celebrationsRoute{i, l, m}

	handler.POST("/celebrations", c.create)
	handler.DELETE("/celebrations/:celebrationID", c.delete)
	handler.GET("/celebrations", c.get)

}

type createCelebrationResponse struct {
	celebration dtos.Celebration `json:"celebration"`
}

type createCelebrationRequest struct {
	Message string `json:"message"  binding:"required"  example:"What is happening?"`
	Media   []struct {
		URL           string `json:"url"  binding:"required"`
		AlternateText string `json:"alternate_text"`
	} `json:"media"`
}

type getCelebrationRequest struct {
	Page  int
	Limit int
}

type getCelebrationResponse struct {
	Page         int                `json:"page"`
	Limit        int                `json:"limit"`
	Celebrations []dtos.Celebration `json:"celebrations"`
}

// @Summary     Create Celebration
// @Description Create a celebration
// @ID          create-celebration
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       request body     createCelebrationRequest true "create a new celebration"
// @Success     201     {object} createCelebrationResponse
// @Router      /celebrations [post]
func (r *celebrationsRoute) create(c *gin.Context) {
	var request createCelebrationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()
	celebration := &models.Celebration{
		Message: &request.Message,
	}

	err := r.uc.Create(ctx, celebration)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		HTTPError(c, http.StatusInternalServerError, "unable to create a celebration")

		return
	}

	c.JSON(http.StatusCreated, createCelebrationResponse{})
}

// @Summary     Get Celebrations
// @Description Get user's celebrations
// @ID          get-celebrations
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       request query    getCelebrationRequest true "create a new celebration"
// @Success     200     {object} getCelebrationResponse
// @Router      /celebrations [get]
func (r *celebrationsRoute) get(c *gin.Context) {
	var request getCelebrationRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	//ctx := c.Request.Context()
	//
	//
	//err := r.uc.Create(ctx, celebration)
	//if err != nil {
	//	r.logger.Error(err, "http - v1 - registerUser")
	//	HTTPError(c, http.StatusInternalServerError, "unable to create a celebration")
	//
	//	return
	//}

	c.JSON(http.StatusCreated, getCelebrationResponse{})
}

type deleteCelebrationResponse struct {
	Message string `json:"message"`
}

// @Summary     Delete Celebration
// @Description Delete a celebration
// @ID          delete-celebration
// @Tags        Celebrations
// @Accept      json
// @Produce     json
// @Param       celebrationID path     string true "celebration post identifier"
// @Success     200           {object} deleteCelebrationResponse
// @Failure     400           {object} ErrorResponse
// @Failure     401           {object} ErrorResponse
// @Failure     404           {object} ErrorResponse
// @Failure     500           {object} ErrorResponse
// @Router      /celebrations/{celebrationID} [delete]
func (r *celebrationsRoute) delete(c *gin.Context) {
	cID := c.Param("celebrationID")

	celebration, err := r.uc.Get(c.Request.Context(), cID)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		HTTPError(c, http.StatusNotFound, "unable to delete celebration")

		return
	}

	// TODO: VALIDATE USER HAS RIGHT TO DELETE CELEBRATION

	err = r.uc.Delete(c.Request.Context(), celebration)
	if err != nil {
		r.logger.Error(err, "http - v1 - celebrations - create")
		HTTPError(c, http.StatusInternalServerError, "unable to delete celebration")

		return
	}

	c.JSON(http.StatusOK, deleteCelebrationResponse{
		Message: "celebration has been successfully deleted",
	})
}
