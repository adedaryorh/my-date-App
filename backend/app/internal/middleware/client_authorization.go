package middleware

import (
	"net/http"

	"backend.app/internal/usecase"
	"backend.app/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

const AuthorizationHeaderKey = "X-Auth-Token"

var allowedPaths = []string{
	"/v1/client/auth",
	"/v1/client",
	"/swagger/index.html",
}

var requiredPaths = []string{
	"/v1/business/industry",
	"/v1/login",
	"/v1/register/initialise",
	"/v1/register/validate",
	"/v1/otp/resend",
}

// ClientAuthorization - ensures that the request comes from an authorised client
func ClientAuthorization(uc usecase.ClientUseCase, l logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(allowedPaths, c.Request.URL.Path) {
			return
		}

		if !slices.Contains(requiredPaths, c.Request.URL.Path) {
			return
		}

		headers := c.Request.Header
		token, exist := headers[AuthorizationHeaderKey]
		if !exist || len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
				Message string
			}{
				Message: "invalid credentials: client authorisation failed",
			})

			return
		}

		err := uc.ValidateClientToken(c, token[0])
		if err != nil {
			l.Debug(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
				Message string
			}{
				Message: "invalid credentials: client authorisation failed",
			})
		}

	}
}
