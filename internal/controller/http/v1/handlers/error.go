package handlers

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse -
type ErrorResponse struct {
	Error string `json:"error" example:"message"`
}

func HTTPError(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, ErrorResponse{msg})
}
