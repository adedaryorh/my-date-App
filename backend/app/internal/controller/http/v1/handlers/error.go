package handlers

import (
	"backend.app/internal/controller/response"
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

func HandleServiceErrorResponse(c *gin.Context, err error, msg string, statusCode int) {
	sErr, ok := err.(*response.ServiceErrorResponse)
	if !ok {
		HTTPErrorWithInformation(c, statusCode, msg, err)
	}

	HTTPErrorWithInformation(c, sErr.StatusCode, msg, sErr.Err)
}
