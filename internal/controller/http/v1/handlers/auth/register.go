package auth

import (
	"celebut-api/internal/contexts"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/services/mailer"
	"celebut-api/internal/services/token"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type registerRoute struct {
	user   usecase.User
	token  token.TokenService
	logger logger.Interface
	mapper mappers.UserMapper
	mailer mailer.Mailer
}

func NewRegisterRoutes(
	handler *gin.RouterGroup,
	i usecase.User,
	l logger.Interface,
	t token.TokenService,
	m mappers.UserMapper,
	ma mailer.Mailer) {
	r := &registerRoute{i, t, l, m, ma}

	handler.POST("/register/initialise", r.initRegister)
	handler.POST("/register/validate", r.validateOTP)
	handler.POST("/register/user", r.completeUserRegistration)
	handler.POST("/register/business", r.completeBusinessRegistration)
}

type initialiseRegisterRequest struct {
	Email       *string `json:"email"              binding:"omitempty,email"  example:"user@email.com"`
	PhoneNumber *string `json:"phone_number"       binding:"omitempty" example:"0712345678"`
	AccountType int     `json:"account_type"       binding:"required"  example:"1"`
}

type initialiseRegisterResponse struct {
	Message string `json:"message"`
}

// @Summary     Initialise Registration
// @Description Begin registration process
// @ID          initialise-register
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request      body     initialiseRegisterRequest true "begin a new registration"
// @Param       x-auth-token header   string                    true "Authorization Token"
// @Success     201          {object} initialiseRegisterResponse
// @Router      /register/initialise [post]
func (r *registerRoute) initRegister(c *gin.Context) {
	var request initialiseRegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - initialise registration")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Email == nil && request.PhoneNumber == nil {
		r.logger.Error("client failed to send in user's email or phone number")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.user.UserIsEnabled(c, "email", *request.Email)
	if err != nil {
		r.logger.Error(err, "http - v1 - initialise registration")
		handlers.HTTPError(c, http.StatusBadRequest, "user already exists")

		return
	}

	// Currently, we are using email for now. Phone number comes in later
	if request.Email != nil {
		err = r.beginRegistrationForUser(c, *request.Email, "email", request.AccountType)
	} else {
		err = r.beginRegistrationForUser(c, *request.PhoneNumber, "phone_number", request.AccountType)
	}

	if err != nil {
		r.logger.Error(err, "http - v1 - initialise registration")
		handlers.HTTPError(c, http.StatusInternalServerError, "internal server error")

		return
	}

	c.JSON(http.StatusOK, initialiseRegisterResponse{
		Message: "A message has been sent to the user via his provided information.",
	})

}

type validateOTPRequest struct {
	Email       *string `json:"email"              binding:"email"  example:"user@email.com"`
	PhoneNumber *string `json:"phone_number"       example:"0712345678"`
	OTP         string  `json:"otp"       binding:"required"  example:"123456"`
}

type validateOTPResponse struct {
	Token string `json:"token"`
}

// @Summary     Validate Registration OTP
// @Description Validate registration OTP
// @ID          validate-register-otp
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request      body     validateOTPRequest true "otp validation body"
// @Param       x-auth-token header   string             true "Authorization Token"
// @Success     200          {object} validateOTPResponse
// @Router      /register/validate [post]
func (r *registerRoute) validateOTP(c *gin.Context) {
	var request validateOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - validate OTP")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Email == nil && request.PhoneNumber == nil {
		r.logger.Error("client failed to send in user's email or phone number")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	user, err := r.user.UserByEmail(c, *request.Email)
	if err != nil {
		r.logger.Error(err, "http - v1 - validate OTP")
		handlers.HTTPError(c, http.StatusBadRequest, "user already exists")

		return
	}

	if user == nil {
		r.logger.Error(err, "http - v1 - validate OTP")
		handlers.HTTPError(c, http.StatusBadRequest, "user doesn't exist")

		return
	}

	if err = r.user.ValidateOTP(c, user.ID, request.OTP, "register"); err != nil {
		r.logger.Error(err, "http - v1 - validate OTP")
		handlers.HTTPError(c, http.StatusBadRequest, "unable to validate OTP")

		return
	}

	claims := token.Claims{
		UserID: user.UserID,
		Email:  user.Email,
	}

	t, err := r.token.GenerateToken(claims)
	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to generate user token")

		return
	}

	c.JSON(http.StatusOK, validateOTPResponse{
		Token: t,
	})

}

type registerResponse struct {
	User dtos.User `json:"user"`
}

type registerBusinessResponse struct {
	User dtos.Business `json:"user"`
}

type registerUserRequest struct {
	FirstName          string  `json:"first_name"  binding:"required"  example:"John"`
	LastName           string  `json:"last_name"  binding:"required"  example:"Doe"`
	UserName           string  `json:"username"  binding:"required"  example:"johndoe"`
	Email              *string `json:"email"       binding:"omitempty,email"  example:"user@email.com"`
	Password           string  `json:"password"       binding:"required"  example:"password"`
	CountryCode        *string `json:"country_code"       binding:"omitempty"  example:"234"`
	PhoneNumber        *string `json:"phone_number"       binding:"omitempty"  example:"0712345678"`
	RelationshipStatus *string `json:"relationship_status" binding:"omitempty"  example:"single"`
	DateOfBirth        string  `json:"dob" binding:"required" example:"2020-10-12"`
	Gender             *string `json:"gender" binding:""  example:"male"`
	ProfileImageURL    *string `json:"profile_image_url" binding:"omitempty,url"  example:"https://abc.png"`
	Interests          *string `json:"interests" binding:""  example:"fashion,entertainment"`
}

type registerBusinessRequest struct {
	BusinessName string  `json:"business_name"  binding:"required"  example:"John XYZ Plc"`
	Email        *string `json:"email"       binding:"email"  example:"user@email.com"`
	Password     string  `json:"password"       binding:"required"  example:"password"`
	CountryCode  *string `json:"country_code"       binding:""  example:"234"`
	PhoneNumber  *string `json:"phone_number"        example:"0712345678"`
	IndustryType int     `json:"industry_type"       binding:"required"  example:"1"`
}

// @Summary     Complete User Registration
// @Description Complete a user's registration
// @ID          register-user
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request       body     registerUserRequest true "complete user registration"
// @Param       x-auth-token  header   string              true "Client authorization token"
// @Param       Authorization header   string              true "Bearer Authorization Token"
// @Success     200           {object} registerResponse
// @Router      /register/user [post]
func (r *registerRoute) completeUserRegistration(c *gin.Context) {
	var request registerUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - register")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	layout := "2006-01-02"
	dobTime, err := time.Parse(layout, request.DateOfBirth)

	if err != nil {
		r.logger.Error(err.Error())
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	ctx := c.Request.Context()

	u, ok := c.Get(contexts.ContextUser)
	if !ok || u == nil {
		handlers.HTTPError(c, http.StatusBadRequest, "unable to get user information")

		return
	}

	user, ok := u.(*models.User)
	if !ok {
		handlers.HTTPError(c, http.StatusBadRequest, "unable to get user information")

		return
	}

	user.FirstName = &request.FirstName
	user.LastName = &request.LastName
	user.Username = &request.UserName
	user.Password = request.Password
	user.DateOfBirth = &dobTime

	if request.Email != nil {
		user.Email = *request.Email
	}

	if request.CountryCode != nil {
		user.CountryCode = request.CountryCode
	}

	if request.PhoneNumber != nil {
		user.PhoneNumber = request.PhoneNumber
	}

	if request.RelationshipStatus != nil {
		user.RelationshipStatus = request.RelationshipStatus
	}

	if request.Gender != nil {
		user.Gender = request.Gender
	}

	if request.Interests != nil {
		user.Interests = request.Interests
	}

	err = r.user.CompleteRegistration(ctx, user)
	if err != nil {
		r.logger.Error(err, "http - v1 - completeUserRegistration")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to register user")

		return
	}

	c.JSON(http.StatusOK, registerResponse{
		User: r.mapper.MapToUserDto(*user),
	})
}

// @Summary     Complete business registration
// @Description Complete business registration
// @ID          register-business
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request       body     registerBusinessRequest true "Complete a new business registration"
// @Param       x-auth-token  header   string                  true "Authorization Token"
// @Param       Authorization header   string                  true "Bearer Authorization Token"
// @Success     200           {object} registerBusinessResponse
// @Router      /register/business [post]
func (r *registerRoute) completeBusinessRegistration(c *gin.Context) {
	var request registerBusinessRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - register")
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	ctx := c.Request.Context()
	u, ok := c.Get(contexts.ContextUser)
	if !ok || u == nil {
		handlers.HTTPError(c, http.StatusBadRequest, "unable to get user information")

		return
	}

	user, ok := u.(models.User)
	if !ok || u == nil {
		handlers.HTTPError(c, http.StatusBadRequest, "unable to get user information")

		return
	}

	user.BusinessName = &request.BusinessName
	user.Password = request.Password
	user.Industry = &models.Industry{
		ID: request.IndustryType,
	}

	if request.Email != nil {
		user.Email = *request.Email
	}

	if request.CountryCode != nil {
		user.CountryCode = request.CountryCode
	}

	if request.PhoneNumber != nil {
		user.PhoneNumber = request.PhoneNumber
	}

	err := r.user.CompleteRegistration(ctx, &user)
	if err != nil {
		r.logger.Error(err, "http - v1 - completeBusinessRegistration")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to register business")

		return
	}

	c.JSON(http.StatusOK, registerBusinessResponse{
		User: r.mapper.MapToBusinessDto(user),
	})
}

// beginRegistrationForUser -
func (r *registerRoute) beginRegistrationForUser(
	c *gin.Context,
	userInfo string,
	infoType string,
	accountType int) error {
	ctx := c.Request.Context()
	//TODO: Sanitise phone number before validation

	otp, err := r.user.RegisterUserWithOTP(ctx, userInfo, infoType, accountType)
	if err != nil {
		return err
	}

	if infoType == "email" {
		err = r.mailer.SendOTP(userInfo, otp.OTP)
	}

	if err != nil {
		return err
	}

	return nil
}
