package handlers

import (
	"celebut-api/internal/dtos"
	"celebut-api/internal/mappers"
	userservice "celebut-api/internal/services/users"
	"celebut-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"time"
)

type usersRoute struct {
	user   userservice.User
	logger logger.Interface
	mapper mappers.UserMapper
}

func NewUserRoutes(
	handler *gin.RouterGroup,
	u userservice.User,
	l logger.Interface,
	m mappers.UserMapper) {

	r := &usersRoute{u, l, m}

	handler.PUT("/users", r.editUser)
	handler.GET("/users/:userID", r.getUser)
}

type editUserResponse struct {
	Status string    `json:"status"`
	User   dtos.User `json:"user"`
}

type editUserRequest struct {
	FirstName    string                `form:"first_name" json:"first_name"  binding:"required"  example:"John"`
	LastName     string                `form:"last_name" json:"last_name"  binding:"required"  example:"Doe"`
	UserName     string                `form:"username" json:"username"  binding:"required"  example:"johndoe"`
	Email        *string               `form:"email" json:"email"       binding:"omitempty,email"  example:"user@email.com"`
	Password     string                `form:"password" json:"password"       binding:"required"  example:"password"`
	CountryCode  *string               `form:"country_code" json:"country_code"       binding:"omitempty"  example:"234"`
	PhoneNumber  *string               `form:"phone_number" json:"phone_number"      binding:"omitempty"  example:"0712345678"`
	DateOfBirth  string                `form:"dob" json:"dob" binding:"required" example:"2020-10-12"`
	ProfileImage *multipart.FileHeader `form:"profile_image" binding:"omitempty" swaggerignore:"true"`
}

// @Summary     Edit user info
// @Description Edit user's profile
// @ID          edit-user
// @Tags        Users
// @Accept      multipart/form-data
// @Produce     json
// @Param       profile_image formData file                false "profile image file"
// @Param       request       formData editUserRequest true "edit user information"
// @Success     200           {object} editUserResponse
// @Security Bearer
// @Router      /user [put]
func (ur *usersRoute) editUser(c *gin.Context) {
	var request editUserRequest

	if err := c.ShouldBind(&request); err != nil {
		ur.logger.Error(err, "http - v1 - user edit")
		HTTPError(c, http.StatusBadRequest, "invalid request body")

		return
	}

	ctx := c.Request.Context()

	layout := "2006-01-02"
	dobTime, err := time.Parse(layout, request.DateOfBirth)

	if err != nil {
		ur.logger.Error(err.Error())
		HTTPError(c, http.StatusBadRequest, err.Error())

		return
	}

	user, err := GetSessionUser(c)
	if err != nil {
		HTTPError(c, http.StatusBadRequest, "unable to get user information")

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

	user.ProfileImage = request.ProfileImage

	err = ur.user.EditUser(ctx, user)
	if err != nil {
		ur.logger.Error(err, "http - v1 - editUser")
		HTTPError(c, http.StatusInternalServerError, "unable to edit user")

		return
	}

	c.JSON(http.StatusOK, editUserResponse{
		Status: "success",
		User:   ur.mapper.MapToUserDto(*user),
	})
}

type getUserResponse struct {
	Status string
	Data   dtos.UserProfile
}

// @Summary     Get user
// @Description Get user profile info
// @ID          get-user-info
// @Tags        Users
// @Accept      json
// @Produce     json
// @Param       userID path     string true "user identifier"
// @Success     200           {object} getUserResponse
// @Failure     400           {object} handlers.ErrorResponse
// @Failure     401           {object} handlers.ErrorResponse
// @Failure     404           {object} handlers.ErrorResponse
// @Failure     500           {object} handlers.ErrorResponse
// @Security    Bearer
// @Router      /users/{userID} [get]
func (ur *usersRoute) getUser(c *gin.Context) {
	userID := c.Param("userID")

	ctx := c.Request.Context()
	user, err := ur.user.GetUser(ctx, userID)
	if err != nil {
		ur.logger.Error(err, "http - v1 - users - getUser")
		HTTPErrorWithInformation(c, http.StatusInternalServerError, "unable to get user info", err)

		return
	}

	c.JSON(http.StatusOK, getUserResponse{
		Status: "success",
		Data:   *user,
	})
}
