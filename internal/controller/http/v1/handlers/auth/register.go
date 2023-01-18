package auth

import (
	"celebut-api/internal/contexts"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/services/file"
	"celebut-api/internal/services/mailer"
	"celebut-api/internal/services/token"
	"celebut-api/internal/usecase"
	"celebut-api/internal/validators"
	"celebut-api/pkg/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type registerRoute struct {
	user         usecase.User
	token        token.TokenService
	logger       logger.Interface
	mapper       mappers.UserMapper
	mailer       mailer.Mailer
	uploadClient file.UploadFileClient
}

func NewRegisterRoutes(
	handler *gin.RouterGroup,
	i usecase.User,
	l logger.Interface,
	t token.TokenService,
	m mappers.UserMapper,
	ma mailer.Mailer,
	uc file.UploadFileClient) {
	r := &registerRoute{i, t, l, m, ma, uc}

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
// @Success     201          {object} initialiseRegisterResponse
// @Security Auth-Token
// @Router      /register/initialise [post]
func (r *registerRoute) initRegister(c *gin.Context) {
	//TODO: Clean up
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

	var err error

	// Currently, we are using email for now. Phone number comes in later
	if request.Email != nil {
		err := r.user.UserIsEnabled(c, "email", *request.Email)
		if err != nil {
			r.logger.Error(err, "http - v1 - initialise registration")
			handlers.HTTPError(c, http.StatusBadRequest, "user already exists")

			return
		}

		err = r.beginRegistrationForUser(c, *request.Email, "email", request.AccountType)
	} else {
		err := r.user.UserIsEnabled(c, "phone", *request.PhoneNumber)
		if err != nil {
			r.logger.Error(err, "http - v1 - initialise registration")
			handlers.HTTPError(c, http.StatusBadRequest, "user already exists")

			return
		}

		err = r.beginRegistrationForUser(c, *request.PhoneNumber, "phone", request.AccountType)
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
	Email       *string `json:"email"         binding:"email,omitempty"  example:"user@email.com"`
	PhoneNumber *string `json:"phone_number"  binding:"omitempty" example:"0712345678"`
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
// @Success     200          {object} validateOTPResponse
// @Security Auth-Token
// @Router      /register/validate [post]
func (r *registerRoute) validateOTP(c *gin.Context) {
	var request validateOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - validate OTP")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "invalid request body", err)

		return
	}

	if request.Email == nil && request.PhoneNumber == nil {
		r.logger.Error("client failed to send in user's email or phone number")
		handlers.HTTPErrorWithInformation(c, http.StatusBadRequest, "invalid request body", errors.New("email/phone number is required"))

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
	FirstName          string                `form:"first_name" json:"first_name"  binding:"required"  example:"John"`
	LastName           string                `form:"last_name" json:"last_name"  binding:"required"  example:"Doe"`
	UserName           string                `form:"username" json:"username"  binding:"required"  example:"johndoe"`
	Email              *string               `form:"email" json:"email"       binding:"omitempty,email"  example:"user@email.com"`
	Password           string                `form:"password" json:"password"       binding:"required"  example:"password"`
	CountryCode        *string               `form:"country_code" json:"country_code"       binding:"omitempty"  example:"234"`
	PhoneNumber        *string               `form:"phone_number" json:"phone_number"      binding:"omitempty"  example:"0712345678"`
	RelationshipStatus *string               `form:"relationship_status" json:"relationship_status" binding:"omitempty"  example:"single"`
	DateOfBirth        string                `form:"dob" json:"dob" binding:"required" example:"2020-10-12"`
	Gender             *string               `form:"gender" json:"gender" binding:""  example:"male"`
	ProfileImage       *multipart.FileHeader `form:"profile_image" binding:"omitempty" swaggerignore:"true"`
	Interests          *string               `form:"interests" json:"interests"  binding:""  example:"fashion,entertainment"`
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
// @Accept      multipart/form-data
// @Produce     json
// @Param       profile_image formData file                false "profile image file"
// @Param       request       formData registerUserRequest true "complete user registration"
// @Success     200           {object} registerResponse
// @Security Bearer
// @Router      /register/user [post]
func (r *registerRoute) completeUserRegistration(c *gin.Context) {
	var request registerUserRequest

	if err := c.ShouldBind(&request); err != nil {
		r.logger.Error(err, "http - v1 - register")
		handlers.HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()

	var profileImageURL *string
	if request.ProfileImage != nil {
		imgFile, err := request.ProfileImage.Open()
		defer imgFile.Close()

		if err != nil {
			r.logger.Error(err, "unable to open file")
			handlers.HTTPError(c, http.StatusBadRequest, "unable to upload file")

			return
		}

		// We only accept PNG, JPEG and JPG for profile images for now...
		_, err = validators.ValidateFileExtension(imgFile, []string{"image/png", "image/jpeg", "image/jpg"})
		if err != nil {
			r.logger.Error(err)
			handlers.HTTPError(c, http.StatusBadRequest, "unable to upload file")

			return
		}

		fileName, err := generateRandomFileName()
		if err != nil {
			r.logger.Error(err, "unable to generate file name")
			handlers.HTTPError(c, http.StatusInternalServerError, "unable to upload file")

			return
		}

		ext := filepath.Ext(request.ProfileImage.Filename)
		fileName = fmt.Sprintf("profile-images/%s.%s", fileName, ext)
		profileImageURL, err = r.uploadClient.UploadToBucket(ctx, "celebut", fileName, imgFile)
	}

	layout := "2006-01-02"
	dobTime, err := time.Parse(layout, request.DateOfBirth)

	if err != nil {
		r.logger.Error(err.Error())
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

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

	if profileImageURL != nil {
		user.ProfileImageURL = profileImageURL
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

// TODO: Move all of this to a user service or file service
func generateRandomFileName() (string, error) {
	newUUID, err := uuid.DefaultGenerator.NewV4()

	if err != nil {
		return "", fmt.Errorf("unable to generate UUID: %s", err)
	}

	fileName := strings.ReplaceAll(newUUID.String(), "-", "")

	return fileName, nil
}

// @Summary     Complete business registration
// @Description Complete business registration
// @ID          register-business
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request       body     registerBusinessRequest true "Complete a new business registration"
// @Success     200           {object} registerBusinessResponse
// @Security Bearer
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
