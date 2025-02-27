package dtos

import "backend.app/common/constants"

// ResponseObject structure for the response object
type ResponseObject struct {
	Code    int                  `json:"-"`
	Status  constants.HttpStatus `json:"status,omitempty"`
	Message string               `json:"message,omitempty"`
	Data    interface{}          `json:"data,omitempty"`
	Error   interface{}          `json:"error,omitempty"`
}
