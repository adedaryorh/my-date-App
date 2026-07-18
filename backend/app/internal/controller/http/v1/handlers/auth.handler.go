package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"backend.app/common/constants"
	"backend.app/common/helpers"
	"backend.app/internal/dtos"
	"backend.app/internal/models"
	"backend.app/pkg/response"
	"github.com/gin-gonic/gin"
)

// @Tags Auth
// @Summary Signs user up
// @Schemes
// @Description Signs user up
// @Param   request   body     dtos.UserSignUp   true  "user's sign up data object"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Router /auth/sign-up/user [post]
func (h *Handler) SignUpUser(c *gin.Context) {
	var input dtos.UserSignUp
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	input.Email = strings.ToLower(input.Email)
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	// send to core
	result := h.core.SignUpUser(c, &input)
	c.JSON(result.Code, result)
}

// @Tags Auth
// @Summary Signs Business up
// @Schemes
// @Description Signs Business up
// @Param   request   body     dtos.BusinessSignUp   true  "business' sign up data object"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject{} "desc"
// @Router /auth/sign-up/business [post]
func (h *Handler) SignUpBusiness(c *gin.Context) {
	var input dtos.BusinessSignUp
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	input.Email = strings.ToLower(input.Email)
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	// send to core
	result := h.core.SignUpBusiness(c, &input)
	c.JSON(result.Code, result)
}

// @Tags Auth
// @Summary Signs user in
// @Schemes
// @Description Signs admin user in
// @Param   request   body     models.SignInDto   true  "user's email and password"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject{data=models.AuthenticatedUser} "desc"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var input models.SignInDto
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	input.Email = strings.ToLower(input.Email)
	if err := helpers.ValidateInput(input); err != nil {
		fmt.Println(err)
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	// send to core
	result := h.core.Login(c, input)
	c.JSON(result.Code, result)
}

// @Tags Auth
// @Summary Confirm User Phone
// @Schemes
// @Description Confirm User Phone
// @Param   request   body     dtos.ConfirmPhoneNumber  true  "data to confirm user's phone"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "desc"
// @Router /auth/confirm-phone [patch]
func (h *Handler) ConfirmPhone(c *gin.Context) {
	var input dtos.ConfirmPhoneNumber
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	// send to controller
	result := h.core.ConfirmPhone(c, &input)
	c.JSON(result.Code, result)
}

func (h *Handler) ConfirmEmail(c *gin.Context) {
	var input dtos.ConfirmEmail
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	input.Email = strings.ToLower(strings.TrimSpace(input.Email))
	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	result := h.core.ConfirmEmail(c, &input)
	c.JSON(result.Code, result)
}

// @Tags Auth
// @Summary Send Password Reset Token
// @Schemes
// @Description Sends User Password Reset Token
// @Param   request   body     dtos.Phone  true  "data to send user password reset link"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "desc"
// @Router /auth/reset-password [post]
func (h *Handler) SendResetPasswordToken(c *gin.Context) {
	var input dtos.Phone
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	// send to controller
	result := h.core.SendResetPasswordToken(c, input.Phone)
	c.JSON(result.Code, result)
}

// @Tags Auth
// @Summary Password Reset
// @Schemes
// @Description Resets User Password
// @Param   request   body     dtos.ResetPassword   true  "data to reset user's password"
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "desc"
// @Router /auth/reset-password [patch]
func (h *Handler) ResetPassword(c *gin.Context) {
	var input dtos.ResetPassword
	// bind input
	if err := c.BindJSON(&input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}

	if err := helpers.ValidateInput(input); err != nil {
		result := response.BadRequestResponse(err, constants.HttpStatusBadRequest)
		c.JSON(result.Code, result)
		return
	}
	// send to controller
	result := h.core.ResetPassword(c, &input)
	c.JSON(result.Code, result)
}

// @Tags Auth
// @Summary Gets Logged In user
// @Schemes
// @Description Gets Logged In user
// @Accept json
// @Produce json
// @Success 200 {object} dtos.ResponseObject "desc"
// @Router /auth/self [get]
func (h *Handler) Me(c *gin.Context) {
	user := c.MustGet("authUser").(models.User)
	c.JSON(http.StatusOK, gin.H{
		"status":  constants.HttpStatusSuccess,
		"data":    user,
		"message": constants.UserSuccessFullyFetched,
	})
}
