package errors

import (
	"encoding/json"
)

type GooglyError struct {
	Status  int
	Message string
	Err     error
}

func (e *GooglyError) Error() string {
	data, _ := json.MarshalIndent(e.Message, "", "  ")
	return string(data)
}

func NewInternalError(err error) *GooglyError {
	if err == nil {
		return nil
	}
	return &GooglyError{
		Status:  500,
		Message: "internal error",
		Err:     err,
	}
}

func NewInvalidRequestError() *GooglyError {

	return &GooglyError{
		Status:  400,
		Message: "Invalid Request",
		Err:     nil,
	}
}
