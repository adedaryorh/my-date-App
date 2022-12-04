package handlers

import (
	"celebut-api/internal/models"
	"celebut-api/internal/usecase"
	"celebut-api/internal/usecase/otp"
	"celebut-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type otpRoute struct {
	user       usecase.User
	otpUseCase otp.UserOTP
	logger     logger.Interface
}

func NewOTPRoute(handler *gin.RouterGroup, i usecase.User, o otp.UserOTP, l logger.Interface) {
	r := &otpRoute{i, o, l}

	handler.POST("/otp/resend", r.resendOTP)

}

type resendOTPResponse struct {
	Message string `json:"message"`
}

type resendOTPRequest struct {
	Email       *string `json:"email"              binding:"omitempty,email"  example:"user@email.com"`
	PhoneNumber *string `json:"phone_number"       binding:"omitempty" example:"0712345678"`
	Mode        string  `json:"mode"       binding:"required"  example:"register"`
}

// @Summary     Resend OTP
// @Description Resend OTP code
// @ID          resend-otp
// @Tags        OTP
// @Accept      json
// @Produce     json
// @Param       request      body     resendOTPRequest true "Resend OTP"
// @Success     200          {object} resendOTPResponse
// @Router      /otp/resend [post]
func (r *otpRoute) resendOTP(c *gin.Context) {
	var request resendOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - resend OTP")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Email == nil && request.PhoneNumber == nil {
		r.logger.Error("client failed to send in user's email or phone number")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	var user *models.User
	var err error

	otpRecipient := otp.Recipient{}

	if request.Email != nil {
		otpRecipient.Info = *request.Email
		otpRecipient.Type = "email"
		user, err = r.user.UserByField(c, "email", *request.Email)
	} else {
		otpRecipient.Info = *request.PhoneNumber
		otpRecipient.Type = "phone"

		user, err = r.user.UserByField(c, "phone", *request.PhoneNumber)
	}

	if err != nil || user == nil {
		r.logger.Error(err, "http - v1 - resend OTP")
		HTTPError(c, http.StatusBadRequest, "user not found")

		return
	}

	err = r.otpUseCase.Create(c.Request.Context(), user.ID, request.Mode, otpRecipient)
	if err != nil {
		r.logger.Error(err, "http - v1 - resend OTP")
		HTTPError(c, http.StatusInternalServerError, "unable to generate OTP")

		return
	}

	c.JSON(http.StatusOK, resendOTPResponse{
		Message: "Success",
	})

}
