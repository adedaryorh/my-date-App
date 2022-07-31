package auth

import (
	"celebut-api/internal/config"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	"celebut-api/internal/models"
	"celebut-api/internal/services"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type registerRoute struct {
	user   usecase.User
	token  services.TokenService
	logger logger.Interface
	mapper mappers.UserMapper
}

func NewRegisterRoute(handler *gin.RouterGroup, i usecase.User, l logger.Interface, t services.TokenService, m mappers.UserMapper) {
	r := &registerRoute{i, t, l, m}

	handler.POST("/register/user", r.registerUser)
	handler.POST("/register/business", r.registerBusiness)

}

type RegisterResponse struct {
	User  dtos.User `json:"user"`
	Token string    `json:"token"`
}

type RegisterBusinessResponse struct {
	User  dtos.Business `json:"user"`
	Token string        `json:"token"`
}

type registerUserRequest struct {
	FirstName          string `json:"first_name"  binding:"required"  example:"John"`
	LastName           string `json:"last_name"  binding:"required"  example:"Doe"`
	UserName           string `json:"username"  binding:"required"  example:"johndoe"`
	Email              string `json:"email"       binding:"required,email"  example:"user@email.com"`
	Password           string `json:"password"       binding:"required"  example:"password"`
	CountryCode        string `json:"country_code"       binding:"required"  example:"234"`
	PhoneNumber        string `json:"phone_number"       binding:"required"  example:"0712345678"`
	RelationshipStatus string `json:"relationship_status" binding:"required"  example:"single"`
	DateOfBirth        string `json:"dob" binding:"required" example:"2020-10-12"`
	Gender             string `json:"gender" binding:"required"  example:"male"`
	Interests          string `json:"interests" binding:"required"  example:"fashion,entertainment"`
}

type registerBusinessRequest struct {
	BusinessName string `json:"business_name"  binding:"required"  example:"John XYZ Plc"`
	Email        string `json:"email"       binding:"required,email"  example:"user@email.com"`
	Password     string `json:"password"       binding:"required"  example:"password"`
	CountryCode  string `json:"country_code"       binding:"required"  example:"234"`
	PhoneNumber  string `json:"phone_number"       binding:"required"  example:"0712345678"`
	IndustryType int    `json:"industry_type"       binding:"required"  example:"1"`
}

// @Summary     User Register
// @Description Register a user
// @ID          register-user
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body     registerUserRequest true "Registers a new user"
// @Success     201     {object} RegisterResponse
// @Router      /register/user [post]
func (r *registerRoute) registerUser(c *gin.Context) {
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
	err = r.user.ValidateExistenceByField(ctx, "username", request.UserName)
	if err != nil {
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	err = r.user.ValidateExistenceByField(ctx, "email", request.Email)
	if err != nil {
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	newUser := &models.User{
		FirstName:          &request.FirstName,
		LastName:           &request.LastName,
		Username:           &request.UserName,
		Email:              request.Email,
		Password:           request.Password,
		CountryCode:        request.CountryCode,
		PhoneNumber:        request.PhoneNumber,
		RelationshipStatus: &request.RelationshipStatus,
		DateOfBirth:        &dobTime,
		Gender:             &request.Gender,
		Interests:          &request.Interests,
		AccountType: models.AccountType{
			ID: config.ACCOUNT_BASIC_ID,
		},
	}

	err = r.user.Register(ctx, newUser)
	if err != nil {
		r.logger.Error(err, "http - v1 - registerUser")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to register user")

		return
	}

	claims := services.Claims{
		Name:  *newUser.Username,
		Email: newUser.Email,
	}

	token, err := r.token.GenerateToken(claims)
	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to login user")

		return
	}

	c.JSON(http.StatusCreated, RegisterResponse{
		User:  r.mapper.MapToUserDto(*newUser),
		Token: token,
	})
}

// @Summary     Business Register
// @Description Register a business
// @ID          register-business
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body     registerBusinessRequest true "Registers a new business"
// @Success     201     {object} RegisterBusinessResponse
// @Router      /register/business [post]
func (r *registerRoute) registerBusiness(c *gin.Context) {
	var request registerBusinessRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - register")
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	ctx := c.Request.Context()
	err := r.user.ValidateExistenceByField(ctx, "email", request.Email)
	if err != nil {
		handlers.HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	newUser := &models.User{
		BusinessName: &request.BusinessName,
		Email:        request.Email,
		Password:     request.Password,
		CountryCode:  request.CountryCode,
		PhoneNumber:  request.PhoneNumber,
		AccountType: models.AccountType{
			ID: config.ACCOUNT_BUSINESS_ID,
		},
		Industry: &models.Industry{
			ID: request.IndustryType,
		},
	}

	err = r.user.Register(ctx, newUser)
	if err != nil {
		r.logger.Error(err, "http - v1 - registerBusiness")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to register business")

		return
	}

	claims := services.Claims{
		Name:  *newUser.BusinessName,
		Email: newUser.Email,
	}

	token, err := r.token.GenerateToken(claims)
	if err != nil {
		r.logger.Error(err, "http - v1 - login")
		handlers.HTTPError(c, http.StatusInternalServerError, "unable to login user")

		return
	}

	c.JSON(http.StatusCreated, RegisterBusinessResponse{
		User:  r.mapper.MapToBusinessDto(*newUser),
		Token: token,
	})
}
