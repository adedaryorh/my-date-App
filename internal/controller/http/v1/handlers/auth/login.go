package auth

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/models/users"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginRoute struct {
	i usecase.Industry
	l logger.Interface
}

func NewLoginRoute(handler *gin.RouterGroup, i usecase.Industry, l logger.Interface) {
	r := &loginRoute{i, l}

	handler.POST("/login", r.login)

}

type loginResponse struct {
	User  users.User `json:"user"`
	Token string     `json:"token"`
}

type loginRequest struct {
	Email       string `json:"email"       binding:"required"  example:"user@email.com"`
	Password    string `json:"password"       binding:"required"  example:"password"`
	AccountType int    `json:"account_type"  binding:"required"  example:"1"`
}

// @Summary     User login
// @Description Login a user
// @ID          login
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       request body registerRequest true "Login user"
// @Success     200 {object} RegisterResponse
// @Router      /login [post]
func (r *loginRoute) login(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - login")
		handlers.ErrorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	//translation, err := r.i.Create(
	//	c.Request.Context(),
	//	business.Industry{
	//		Name:        request.Name,
	//		Description: request.Description,
	//	},
	//)
	//if err != nil {
	//	r.l.Error(err, "http - v1 - createIndustry")
	//	//errorResponse(c, http.StatusInternalServerError, "translation service problems")
	//
	//	return
	//}

	c.JSON(http.StatusOK, loginResponse{})
}
