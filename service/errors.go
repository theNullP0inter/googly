package service

import "encoding/json"

// ServiceError is an error with following properties
// Code: it's synonymous to http status code
// Err: actual error if any
// Errors: Generally used to show validation errors
type ServiceError struct {
	Code    int
	Message string
	Err     error
	Errors  interface{}
}

// Error() is required to be implemented by error interface
func (e *ServiceError) Error() string {
	data, _ := json.MarshalIndent(e.Message, "", "  ")
	return string(data)
}

// NewServiceError creates a new ServiceError
func NewServiceError(code int, message string, err error, errs interface{}) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Err:     nil,
		Errors:  errs,
	}
}

// NewInternalServiceError creates a new ServiceError with Code 500
func NewInternalServiceError(err error) *ServiceError {
	return NewServiceError(
		500, "internal error", err, nil,
	)
}

// NewNotFoundServiceError creates a new ServiceError with Code 404
func NewNotFoundServiceError(err error) *ServiceError {
	return NewServiceError(
		404, "not found", err, nil,
	)
}

// NewBadRequestError creates a new ServiceError with Code 400
func NewBadRequestError(err error) *ServiceError {
	return NewServiceError(
		400, "bad request", err, nil,
	)
}
