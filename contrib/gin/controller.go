package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

type GinControllerIngress interface {
	AddRoutes(*gin.RouterGroup)
}

type GinControllerInterface interface {
	controller.Controller
	HttpResponse(*gin.Context, interface{}, int)
	HttpReplySuccess(*gin.Context, interface{})
	HttpReplyServiceError(*gin.Context, *service.ServiceError)
}

type GinController struct {
	*controller.BaseController
}

func HandleHttpError(c *gin.Context, e *controller.HttpError) {
	message := e.Message
	errors := e.Errors

	if e.Code >= 500 {
		if viper.GetString("GIN_MODE") == "release" {
			message = controller.ErrHttpInternal
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

func (c *GinController) HttpResponse(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}

func (c *GinController) HttpReplySuccess(context *gin.Context, data interface{}) {
	c.HttpResponse(context, gin.H{"data": data}, http.StatusOK)
}

func (c *GinController) HttpReplyGinBindError(context *gin.Context, err error) {
	e := &controller.HttpError{
		Code:    422,
		Message: controller.ErrHttpInvalidRequest,
		Err:     nil,
		Errors:  err.Error(),
	}
	HandleHttpError(context, e)
}

func (c *GinController) HttpReplyGinPathParamError(context *gin.Context, err error) {
	e := &controller.HttpError{
		Code:    400,
		Message: controller.ErrHttpInvalidPathParam,
		Err:     err,
		Errors:  nil,
	}
	HandleHttpError(context, e)
}

func (c *GinController) HttpReplyGinNotFoundError(context *gin.Context, err error) {
	e := &controller.HttpError{
		Code:    400,
		Message: controller.ErrHttpInvalidRequest,
		Err:     nil,
		Errors:  err.Error(),
	}
	HandleHttpError(context, e)
}

func (c *GinController) HttpReplyHttpError(context *gin.Context, err *controller.HttpError) {

	HandleHttpError(context, err)
}

func (c *GinController) HttpReplyServiceError(context *gin.Context, err *service.ServiceError) {
	e := controller.NewHttpErrorFromServiceError(err)
	HandleHttpError(context, e)
}

func NewGinController(logger logger.GooglyLoggerInterface) *GinController {
	con := controller.NewBaseController(logger)
	return &GinController{
		con,
	}
}
