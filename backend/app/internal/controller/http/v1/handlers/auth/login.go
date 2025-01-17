package auth

import (
	"net/http"

	"backend.app/internal/config"
	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/internal/dtos"
	"backend.app/internal/mappers"
	"backend.app/internal/services/token"
	"backend.app/internal/usecase"
	"backend.app/pkg/logger"

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
	Status string    `json:"status"`
	User   dtos.User `json:"user"`
	Token  string    `json:"token"`
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
// @Param       request      body     loginRequest true "Login user"
// @Success     200          {object} loginUserResponse
// @Security Auth-Token
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
		handlers.HandleServiceErrorResponse(c, err, "unable to login user", http.StatusInternalServerError)

		return
	}

	claims := token.Claims{
		Email:  user.Email,
		UserID: user.UserID,
	}

	if request.AccountType == config.ACCOUNT_BASIC_ID {
		claims.Name = *user.Username
	} else {
		claims.Name = *user.BusinessName
	}

	t, err := r.tokenService.GenerateToken(claims)

	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to login user")

		return
	}

	c.JSON(http.StatusOK, loginUserResponse{
		Status: "success",
		User:   r.mapper.MapToUserDto(*user),
		Token:  t,
	})
}
