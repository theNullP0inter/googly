package service

import "encoding/json"

type ServiceError struct {
	Code    int
	Message string
	Err     error
	Errors  interface{}
}

func (e *ServiceError) Error() string {
	data, _ := json.MarshalIndent(e.Message, "", "  ")
	return string(data)
}

func NewServiceError(code int, message string, err error, errs interface{}) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Err:     nil,
		Errors:  errs,
	}
}

func NewInternalServiceError(err error) *ServiceError {
	return NewServiceError(
		500, "internal error", err, nil,
	)
}

func NewNotFoundServiceError(err error) *ServiceError {
	return NewServiceError(
		404, "not found", err, nil,
	)
}

func NewBadRequestError(err error) *ServiceError {
	return NewServiceError(
		400, "bad request", err, nil,
	)
}
