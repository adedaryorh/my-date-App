package middleware

import (
	"celebut-api/internal/contexts"
	"celebut-api/internal/services/token"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
)

const BearerAuthorizationHeaderKey = "Authorization"

var skipAuthPaths = []string{
	"/v1/client/auth",
	"/v1/client",
	"/swagger/index.html",
	"/v1/register/initialise",
	"/v1/register/validate",
	"/v1/login",
}

// Authorization - ensures user has been authenticated.
func Authorization(uc usecase.UserUseCase, t token.TokenService, l logger.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(skipAuthPaths, c.Request.URL.Path) {
			return
		}

		headers := c.Request.Header
		authToken, exist := headers[BearerAuthorizationHeaderKey]
		if !exist || len(authToken) == 0 {
			abortRequest(c)

			return
		}

		bearerToken := strings.Split(authToken[0], " ")
		if len(bearerToken) != 2 {
			abortRequest(c)

			return
		}

		claims, err := t.ValidateToken(bearerToken[1])
		if err != nil {
			abortRequest(c)

			return
		}

		user, err := uc.UserByField(c, "user_id", claims.UserID)
		if err != nil {
			l.Debug(err)
			abortRequest(c)

			return
		}

		c.Set(contexts.ContextUserID, user.UserID)
		c.Set(contexts.ContextUser, user)
	}
}

func abortRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
		Message string
	}{
		Message: "invalid credentials: client authorisation failed",
	})
}
