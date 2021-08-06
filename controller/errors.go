package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/service"
)

type HttpError struct {
	Code    int
	Message string
	Errors  interface{}
	Err     error
}

func (e *HttpError) RespondToGin(c *gin.Context) {
	message := e.Message
	errors := e.Errors

	if e.Code >= 500 {
		if viper.GetString("GIN_MODE") == "release" {
			message = ErrHttpInternal
			errors = nil
		}
	}

	c.JSON(e.Code, gin.H{
		"error": gin.H{
			"message": message,
			"errors":  errors,
		},
	})

}

func NewHttpErrorFromServiceError(err *service.ServiceError) *HttpError {
	return &HttpError{
		Code:    err.Code,
		Message: err.Message,
		Errors:  err.Errors,
		Err:     err.Err,
	}
}
