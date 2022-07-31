package auth

import (
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/models"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clientAuthRoute struct {
	i usecase.ClientUseCase
	l logger.Interface
}

func NewClientAuthRoute(handler *gin.RouterGroup, c usecase.ClientUseCase, l logger.Interface) {
	r := &clientAuthRoute{c, l}

	handler.POST("/client/auth", r.authenticate)
	handler.POST("/client", r.create)

}

type authResponse struct {
	Token string `json:"token"`
}

type authRequest struct {
	ClientID string `json:"client_id"       binding:"required"  example:"abcede"`
	Secret   string `json:"secret"       binding:"required"  example:"password"`
}

// @Summary     Authenticate a client
// @Description Authenticate a client
// @ID          auth-client
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body     authRequest true "Authenticate a client"
// @Success     200     {object} authResponse
// @Router      /client/auth [post]
func (r *clientAuthRoute) authenticate(c *gin.Context) {
	var request authRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 -  client authenticate")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	token, err := r.i.AuthenticateClient(c.Request.Context(), request.ClientID, request.Secret)

	if err != nil {
		r.l.Error(err, "http - v1 - client authenticate")
		handlers.HTTPError(c, http.StatusUnauthorized, "invalid credentials: client authentication failed")
		return
	}

	c.JSON(http.StatusOK, authResponse{
		Token: *token,
	})
}

type createRequest struct {
	Name     string `json:"client_name" binding:"required"  example:"Celebut iOS"`
	ClientID string `json:"client_id"   binding:"required"  example:"abcede"`
	Secret   string `json:"secret"      binding:"required"  example:"password"`
}

type createResponse struct {
	Message string `json:"message"`
}

// @Summary     [Admin] Create a client
// @Description Create a new client
// @ID          create-client
// @Tags        Client
// @Accept      json
// @Produce     json
// @Param       request body     createRequest true "Create a client"
// @Success     200     {object} createResponse
// @Router      /client [post]
func (r *clientAuthRoute) create(c *gin.Context) {
	var request createRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 -  client create")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	client := &models.Client{
		Name:     request.Name,
		ClientID: request.ClientID,
		Secret:   request.Secret,
	}

	err := r.i.CreateClient(c.Request.Context(), client)

	if err != nil {
		r.l.Error(err, "http - v1 - client authenticate")
		handlers.HTTPError(c, http.StatusUnauthorized, "invalid credentials: client authentication failed")
		return
	}

	c.JSON(http.StatusOK, createResponse{
		Message: "client created successfully",
	})
}
