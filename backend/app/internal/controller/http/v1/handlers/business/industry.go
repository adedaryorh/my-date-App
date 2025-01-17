package business

import (
	"net/http"

	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/models"
	"backend.app/internal/usecase"
	"backend.app/pkg/logger"

	"github.com/gin-gonic/gin"
)

type industryRoutes struct {
	i usecase.Industry
	l logger.Interface
	m mappers.IndustryMapper
}

func NewIndustryRoutes(handler *gin.RouterGroup, i usecase.Industry, l logger.Interface, m mappers.IndustryMapper) {
	r := &industryRoutes{i, l, m}

	h := handler.Group("/business/industry")
	{
		h.GET("", r.industries)
		h.POST("", r.createIndustry)
	}
}

type getIndustryResponse struct {
	Status string          `json:"status"`
	Data   []dtos.Industry `json:"industries"`
}

type createIndustryResponse struct {
	Status string        `json:"status"`
	Data   dtos.Industry `json:"industry"`
}

// @Summary     Get industries
// @Description Get list of business industries
// @ID          industries-list
// @Tags        Industries
// @Accept      json
// @Produce     json
// @Success     200          {object} getIndustryResponse
// @Security Auth-Token
// @Router      /business/industry [get]
func (r *industryRoutes) industries(c *gin.Context) {
	industries, err := r.i.Industries(c.Request.Context())

	if err != nil {
		r.l.Error(err, "http - v1 - history")
		handlers.HTTPError(c, http.StatusInternalServerError, "internal server error")

		return
	}

	c.JSON(http.StatusOK, getIndustryResponse{
		Status: "success",
		Data:   r.m.MapToIndustryListDto(industries),
	})
}

type createIndustryRequest struct {
	Name        string `json:"name"       binding:"required"  example:"Engineering"`
	Description string `json:"description"  binding:"required"  example:"description"`
}

// @Summary     Create industry
// @Description Create a new business industry
// @ID          create-industry
// @Tags        Industries
// @Accept      json
// @Produce     json
// @Param       request      body     createIndustryRequest true "Create new industry"
// @Success     200          {object} models.Industry
// @Security Auth-Token
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

	c.JSON(http.StatusCreated, createIndustryResponse{
		Status: "success",
		Data:   r.m.MapToIndustryDto(*industry),
	})
}
