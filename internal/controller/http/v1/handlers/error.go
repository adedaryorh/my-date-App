package handlers

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse -
type ErrorResponse struct {
	Status               string `json:"status" example:"error"`
	Error                string `json:"error" example:"message"`
	DeveloperInformation string `json:"developer_information" example:"an error occurred"`
}

func HTTPError(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, ErrorResponse{"error", msg, ""})
}

func HTTPErrorWithInformation(c *gin.Context, code int, msg string, err error) {
	c.AbortWithStatusJSON(code, ErrorResponse{"error", msg, err.Error()})
}
