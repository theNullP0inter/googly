package service

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ServiceErrorMap map[string][]string

type ServiceError struct {
	Errors       ServiceErrorMap `json:"errors"`
	HttpStatus   int             `json:"http_status"`
	ErrorMessage string          `json:"error_message"`
}

var DefaultServerErrorMap ServiceErrorMap = ServiceErrorMap{"error_message": []string{}}

func (e ServiceError) Error() string {
	data, _ := json.MarshalIndent(e.Errors, "", "  ")
	return string(data)
}

func (e ServiceError) RespondToHttp(c *gin.Context) {
	c.JSON(e.HttpStatus, e)
}

func NewBinIdError() *ServiceError {
	return &ServiceError{
		Errors:       DefaultServerErrorMap,
		HttpStatus:   400,
		ErrorMessage: VALIDATION_FAILED,
	}

}

func NewServiceRequestValidationError(e ServiceErrorMap) *ServiceError {
	return &ServiceError{
		Errors:       e,
		HttpStatus:   422,
		ErrorMessage: VALIDATION_FAILED,
	}

}

func NewInternalServiceError(err error) *ServiceError {
	return &ServiceError{
		Errors:       ServiceErrorMap{"internal_error": []string{err.Error()}},
		HttpStatus:   500,
		ErrorMessage: INTERNAL_ERROR,
	}

}

func NewUniqueConstraintError(resource string) *ServiceError {
	return &ServiceError{
		Errors:       DefaultServerErrorMap,
		HttpStatus:   400,
		ErrorMessage: resource + UNIQUE_CONSTRAINT_ERROR,
	}

}
