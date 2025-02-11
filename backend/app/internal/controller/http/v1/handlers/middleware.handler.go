package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend.app/common/constants"
	"backend.app/common/messages"
)

func (h *Handler) AuthenticatedUserMiddleware() gin.HandlerFunc {
	// add the middleware function
	return func(c *gin.Context) {
		user, err := h.core.Middleware().JwtUserAuth(c)
		if err != nil {
			status := http.StatusUnauthorized
			if err == messages.ErrInvalidToken {
				status = http.StatusGone
			}
			c.JSON(status, gin.H{
				"status":  constants.HttpStatusBadRequest,
				"error":   err,
				"message": messages.ErrInvalidInput.Error(),
			})
			c.Abort()
		} else {
			c.Set("authUser", *user)
		}
		c.Next()
	}
}
