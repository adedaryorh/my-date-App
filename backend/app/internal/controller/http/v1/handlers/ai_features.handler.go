package handlers

import (
	"fmt"
	"strconv"

	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetNearbyCelebrations(c *gin.Context) {
	user := c.MustGet("authUser").(models.User)
	latitude, longitude := user.Latitude, user.Longitude
	var err error
	if raw := c.Query("latitude"); raw != "" {
		latitude, err = strconv.ParseFloat(raw, 64)
		if err != nil {
			result := response.BadRequestResponse(fmt.Errorf("invalid latitude"))
			c.JSON(result.Code, result)
			return
		}
	}
	if raw := c.Query("longitude"); raw != "" {
		longitude, err = strconv.ParseFloat(raw, 64)
		if err != nil {
			result := response.BadRequestResponse(fmt.Errorf("invalid longitude"))
			c.JSON(result.Code, result)
			return
		}
	}
	if latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		result := response.BadRequestResponse(fmt.Errorf("coordinates out of range"))
		c.JSON(result.Code, result)
		return
	}
	radius, _ := strconv.ParseFloat(c.DefaultQuery("radius_km", "50"), 64)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	result := h.core.GetNearbyCelebrations(c.Request.Context(), latitude, longitude, radius, limit)
	c.JSON(result.Code, result)
}

func (h *Handler) GetModerationQueue(c *gin.Context) {
	user := c.MustGet("authUser").(models.User)
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	result := h.core.GetModerationQueue(c.Request.Context(), &user, c.DefaultQuery("status", "pending"), limit)
	c.JSON(result.Code, result)
}

func (h *Handler) ResolveModeration(c *gin.Context) {
	user := c.MustGet("authUser").(models.User)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		result := response.BadRequestResponse(fmt.Errorf("invalid moderation ID"))
		c.JSON(result.Code, result)
		return
	}
	var input dtos.ResolveModerationRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		result := response.BadRequestResponse(err)
		c.JSON(result.Code, result)
		return
	}
	result := h.core.ResolveModeration(c.Request.Context(), &user, id, input.Decision)
	c.JSON(result.Code, result)
}
