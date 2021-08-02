package errors

import "encoding/json"

type GogetaError struct {
	Status  int
	Message string
	Err     error
}

func (e *GogetaError) Error() string {
	data, _ := json.MarshalIndent(e.Message, "", "  ")
	return string(data)
}

func NewInternalError(err error) *GogetaError {
	if err == nil {
		return nil
	}
	return &GogetaError{
		Status:  500,
		Message: "internal error",
		Err:     err,
	}
}
