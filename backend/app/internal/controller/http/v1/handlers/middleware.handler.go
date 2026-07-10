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
			status := constants.HttpStatusTokenExpired
			if err == messages.ErrInvalidToken {
				status = constants.HttpStatusInvalidToken
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  status,
				"error":   err.Error(),
				"message": messages.ErrInvalidInput.Error(),
			})
			c.Abort()
		} else {
			c.Set("authUser", *user)
		}
		c.Next()
	}
}

// @Tags Settings
// @Summary Logs a user out
// @Schemes
// @Description Logs a user out
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "desc"
// @Router /settings/logout [put]
func (h *Handler) LogoutMiddleware() gin.HandlerFunc {
	// add the middleware function
	return func(c *gin.Context) {
		if err := h.core.Middleware().JwtLogUserOut(c); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  constants.HttpStatusBadRequest,
				"error":   err.Error(),
				"message": "logout not successful",
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  constants.HttpStatusSuccess,
			"message": "logout successful",
		})
		c.Abort()
	}

}
