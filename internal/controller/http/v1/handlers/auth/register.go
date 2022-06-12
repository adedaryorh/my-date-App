package auth

import (
	"celebut-api/internal/models/users"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerRoute struct {
	i usecase.Industry
	l logger.Interface
}

func NewRegisterRoute(handler *gin.RouterGroup, i usecase.Industry, l logger.Interface) {
	r := &registerRoute{i, l}

	handler.POST("/register", r.register)

}

type RegisterResponse struct {
	User  users.User `json:"user"`
	Token string     `json:"token"`
}

type registerRequest struct {
	Email       string `json:"email"       binding:"required"  example:"user@email.com"`
	Password    string `json:"password"       binding:"required"  example:"password"`
	AccountType int    `json:"account_type"  binding:"required"  example:"1"`
}

// @Summary     User Register
// @Description Register a user
// @ID          register
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       request body registerRequest true "Registers a new user"
// @Success     201 {object} RegisterResponse
// @Router      /register [post]
func (r *registerRoute) register(c *gin.Context) {
	var request registerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - register")
		//errorResponse(c, http.StatusBadRequest, "invalid request body")

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

	c.JSON(http.StatusOK, RegisterResponse{
	})
}
