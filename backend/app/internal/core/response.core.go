package core

import (
	"net/http"

	"backend.app/common/constants"
	"backend.app/internal/dtos"
)

func successResponse(code int, status constants.HttpStatus, message string, data interface{}) *dtos.ResponseObject {
	obj := dtos.ResponseObject{
		Code:    code,
		Message: message,
		Status:  status,
	}
	if data != nil {
		obj.Data = data
	}
	return &obj

}

func failureResponse(code int, status constants.HttpStatus, message string, err error) *dtos.ResponseObject {
	return &dtos.ResponseObject{
		Code:    code,
		Message: message,
		Error:   err,
		Status:  status,
	}
}

func CreatedSuccessResponse(message string, data interface{}) *dtos.ResponseObject {
	return successResponse(http.StatusCreated, constants.HttpStatusSuccess, message, data)
}

func SuccessResponse(message string, data interface{}) *dtos.ResponseObject {
	return successResponse(http.StatusOK, constants.HttpStatusSuccess, message, data)
}

func ServerErrorResponse(err error, message ...string) *dtos.ResponseObject {
	responseMessage := err.Error()
	if len(message) > 0 {
		responseMessage = message[0]
	}
	return failureResponse(http.StatusInternalServerError, constants.HttpStatusServerError, responseMessage, err)
}
func BadRequestResponse(err error, status ...constants.HttpStatus) *dtos.ResponseObject {
	stat := constants.HttpStatusBadRequest
	if len(status) > 0 {
		stat = status[0]
	}
	return failureResponse(http.StatusBadRequest, stat, err.Error(), err)
}
