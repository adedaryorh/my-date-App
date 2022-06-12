package middleware

import (
	"celebut-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
)

const AuthorizationHeaderKey = "x-auth-token"

var allowedPaths = []string{
	"/v1/client/auth",
	"/v1/client",
	"/swagger/index.html",
}

// Authorization - .
func Authorization(uc usecase.ClientUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(allowedPaths, c.Request.URL.Path) {
			return
		}

		headers := c.Request.Header

		token, exist := headers[AuthorizationHeaderKey]
		if !exist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
				Message string
			}{
				Message: "invalid credentials: client authorisation failed",
			})
		}

		err := uc.ValidateClientToken(c, token[0])
		if err != nil {
			//handlers.ErrorResponse(c, http.StatusUnauthorized, "invalid credentials: client authentication failed")
			c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
				Message string
			}{
				Message: "invalid credentials: client authorisation failed",
			})
		}

	}
}
