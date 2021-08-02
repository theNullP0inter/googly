package errors

import "fmt"

func NewBinIdAssertionError(id interface{}) *GooglyError {
	return &GooglyError{
		Status:  400,
		Message: fmt.Sprintf("%s is Not a valid UUID", id),
		Err:     nil,
	}
}
