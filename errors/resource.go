package errors

import "fmt"

func NewUniqueConstraintError(resource_name string, err error) *GogetaError {
	return &GogetaError{
		Status:  400,
		Message: fmt.Sprintf("%s already exists", resource_name),
		Err:     err,
	}
}

func NewResourceNotFoundError(resource_name string, err error) *GogetaError {
	return &GogetaError{
		Status:  404,
		Message: fmt.Sprintf("%s not found", resource_name),
		Err:     err,
	}
}
