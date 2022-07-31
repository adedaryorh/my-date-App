package middleware

import (
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
)

const AuthorizationHeaderKey = "X-Auth-Token"

var allowedPaths = []string{
	"/v1/client/auth",
	"/v1/client",
	"/swagger/index.html",
}

// ClientAuthorization - .
func ClientAuthorization(uc usecase.ClientUseCase, l logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(allowedPaths, c.Request.URL.Path) {
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
