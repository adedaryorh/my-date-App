package auth

import (
	"celebut-api/internal/config"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/services/token"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRoute struct {
	user         usecase.User
	logger       logger.Interface
	mapper       mappers.UserMapper
	tokenService token.TokenService
}

func NewLoginRoute(handler *gin.RouterGroup, i usecase.User, l logger.Interface, m mappers.UserMapper, t token.TokenService) {
	r := &loginRoute{i, l, m, t}

	handler.POST("/login", r.login)

}

type loginUserResponse struct {
	User  dtos.User `json:"user"`
	Token string    `json:"token"`
}

type loginBusinessResponse struct {
	User  dtos.Business `json:"user"`
	Token string        `json:"token"`
}

type loginRequest struct {
	Email       string `json:"email"       binding:"required"  example:"user@email.com"`
	Password    string `json:"password"       binding:"required"  example:"password"`
	AccountType int    `json:"account_type"  binding:"required"  example:"1"`
}

// @Summary     User login
// @Description Login a user
// @ID          login
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       x-auth-token header   string       true "Authorization Token"
// @Param       request      body     loginRequest true "Login user"
// @Success     200          {object} loginUserResponse
// @Success     200          {object} loginBusinessResponse
// @Router      /login [post]
func (r *loginRoute) login(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()
	user, err := r.user.Login(ctx, request.Email, request.Password, request.AccountType)

	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to login user")

		return
	}

	claims := token.Claims{
		Email: user.Email,
	}

	if request.AccountType == config.ACCOUNT_BASIC_ID {
		claims.Name = *user.Username
	} else {
		claims.Name = *user.BusinessName
	}

	token, err := r.tokenService.GenerateToken(claims)

	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to login user")

		return
	}

	if request.AccountType == config.ACCOUNT_BASIC_ID {
		c.JSON(http.StatusOK, loginUserResponse{
			User:  r.mapper.MapToUserDto(*user),
			Token: token,
		})

		return
	}

	c.JSON(http.StatusOK, loginBusinessResponse{
		User:  r.mapper.MapToBusinessDto(*user),
		Token: token,
	})

}
