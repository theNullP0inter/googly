package controller

import (
	"github.com/theNullP0inter/googly/service"
)

type HttpError struct {
	Code    int
	Message string
	Errors  interface{}
	Err     error
}

func NewHttpErrorFromServiceError(err *service.ServiceError) *HttpError {
	return &HttpError{
		Code:    err.Code,
		Message: err.Message,
		Errors:  err.Errors,
		Err:     err.Err,
	}
}
