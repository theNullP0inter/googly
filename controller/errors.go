package controller

import (
	"github.com/theNullP0inter/googly/service"
)

// HttpError is an http error with message, validation errors, etc
//
// routers controllers can use this to respond errors to requests accordingly
type HttpError struct {
	Code    int
	Message string
	Errors  interface{}
	Err     error
}

// NewHttpErrorFromServiceError creates a new HttpError from a ServiceError
func NewHttpErrorFromServiceError(err *service.ServiceError) *HttpError {
	return &HttpError{
		Code:    err.Code,
		Message: err.Message,
		Errors:  err.Errors,
		Err:     err.Err,
	}
}
