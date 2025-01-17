package middleware

import (
	"errors"
	"net/http"
	"strings"

	"backend.app/internal/contexts"
	"backend.app/internal/services/token"
	"backend.app/internal/usecase"
	"backend.app/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

const BearerAuthorizationHeaderKey = "Authorization"

var skipAuthPaths = []string{
	"/v1/client/auth",
	"/v1/client",
	"/swagger/index.html",
	"/v1/register/initialise",
	"/v1/register/validate",
	"/v1/otp/resend",
	"/v1/login",
	"/v1/business/industry/",
	"/v1/business/industry",
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
			abortRequestWithError(c, errors.New("invalid bearer token"))
			return
		}

		claims, err := t.ValidateToken(bearerToken[1])
		if err != nil {
			abortRequestWithError(c, err)
			return
		}

		user, err := uc.UserByField(c, "user_id", claims.UserID)
		if err != nil {
			l.Debug(err)
			abortRequestWithError(c, err)

			return
		}

		if user == nil {
			abortRequestWithError(c, errors.New("user not found"))

			return
		}

		c.Set(contexts.ContextUserID, user.UserID)
		c.Set(contexts.ContextUser, user)
	}
}

func abortRequest(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "error",
		Message: "invalid credentials: client authorisation failed",
	})
}

func abortRequestWithError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, struct {
		Status               string `json:"status"`
		Message              string `json:"message"`
		DeveloperInformation string `json:"developer_information"`
	}{
		Status:               "error",
		Message:              "invalid credentials: client authorisation failed",
		DeveloperInformation: err.Error(),
	})
}
