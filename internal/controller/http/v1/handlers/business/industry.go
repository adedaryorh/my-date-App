package business

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/models"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type industryRoutes struct {
	i usecase.Industry
	l logger.Interface
}

func NewIndustryRoutes(handler *gin.RouterGroup, i usecase.Industry, l logger.Interface) {
	r := &industryRoutes{i, l}

	h := handler.Group("/business/industry")
	{
		h.GET("/", r.industries)
		h.POST("/", r.createIndustry)
	}
}

type getIndustryResponse struct {
	Industries []models.Industry `json:"industries"`
}

// @Summary     Get industries
// @Description Get list of business industries
// @ID          industries-list
// @Tags        Industries
// @Accept      json
// @Produce     json
// @Param       x-auth-token header   string true "Authorization Token"
// @Success     200          {object} getIndustryResponse
// @Router      /business/industry [get]
func (r *industryRoutes) industries(c *gin.Context) {
	industries, err := r.i.Industries(c.Request.Context())

	if err != nil {
		r.l.Error(err, "http - v1 - history")
		handlers.HTTPError(c, http.StatusInternalServerError, "internal server error")

		return
	}

	c.JSON(http.StatusOK, getIndustryResponse{industries})
}

type createIndustryRequest struct {
	Name        string `json:"name"       binding:"required"  example:"auto"`
	Description string `json:"description"  binding:"required"  example:"en"`
}

// @Summary     Create industry
// @Description Create a new business industry
// @ID          create-industry
// @Tags        Industries
// @Accept      json
// @Produce     json
// @Param       x-auth-token header   string                true "Authorization Token"
// @Param       request      body     createIndustryRequest true "Create new industry"
// @Success     200          {object} models.Industry
// @Router      /business/industry [post]
func (r *industryRoutes) createIndustry(c *gin.Context) {
	var request createIndustryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - createIndustry")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	industry := &models.Industry{
		Name:        request.Name,
		Description: request.Description,
	}

	err := r.i.Create(c.Request.Context(), industry)
	if err != nil {
		r.l.Error(err, "http - v1 - createIndustry")
		handlers.HTTPError(c, http.StatusBadRequest, "internal server error")

		return
	}

	c.JSON(http.StatusOK, industry)
}
