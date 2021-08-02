package errors

import "fmt"

func NewBinIdAssertionError(id interface{}) *GogetaError {
	return &GogetaError{
		Status:  400,
		Message: fmt.Sprintf("%s is Not a valid UUID", id),
		Err:     nil,
	}
}
