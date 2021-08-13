package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

type GinControllerIngress interface {
	AddRoutes(*gin.RouterGroup)
}

type GinControllerInterface interface {
	ControllerInterface
	HttpResponse(*gin.Context, interface{}, int)
	HttpReplySuccess(*gin.Context, interface{})
	HttpReplyServiceError(*gin.Context, *service.ServiceError)
}

type GinController struct {
	*Controller
}

func (c *GinController) HttpResponse(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}

func (c *GinController) HttpReplySuccess(context *gin.Context, data interface{}) {
	c.HttpResponse(context, gin.H{"data": data}, http.StatusOK)
}

func (c *GinController) HttpReplyGinBindError(context *gin.Context, err error) {
	e := &HttpError{
		Code:    422,
		Message: ErrHttpInvalidRequest,
		Err:     nil,
		Errors:  err.Error(),
	}
	e.RespondToGin(context)
}

func (c *GinController) HttpReplyGinPathParamError(context *gin.Context, err error) {
	e := &HttpError{
		Code:    400,
		Message: ErrHttpInvalidPathParam,
		Err:     err,
		Errors:  nil,
	}
	e.RespondToGin(context)
}

func (c *GinController) HttpReplyGinNotFoundError(context *gin.Context, err error) {
	e := &HttpError{
		Code:    400,
		Message: ErrHttpInvalidRequest,
		Err:     nil,
		Errors:  err.Error(),
	}
	e.RespondToGin(context)
}

func (c *GinController) HttpReplyHttpError(context *gin.Context, err HttpError) {
	err.RespondToGin(context)
}

func (c *GinController) HttpReplyServiceError(context *gin.Context, err *service.ServiceError) {
	NewHttpErrorFromServiceError(err).RespondToGin(context)
}

func NewGinController(logger logger.GooglyLoggerInterface) *GinController {
	controller := NewController(logger)
	return &GinController{
		controller,
	}
}
