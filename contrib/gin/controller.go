package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

// GinControllerIngress should be implemented by the controller for it to be able to register its routes
type GinControllerIngress interface {
	AddRoutes(*gin.RouterGroup)
}

// GinController should be implemented to connect gin
type GinController interface {
	controller.Controller
	HttpResponse(*gin.Context, interface{}, int)
	HttpReplySuccess(*gin.Context, interface{})
	HttpReplyServiceError(*gin.Context, *service.ServiceError)
}

// BaseGinController is a basic implementation of GinController
type BaseGinController struct {
	*controller.BaseController
}

// HandleHttpError converts controller.HttpError to a gin error
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

// HttpResponse sends user a http response with given data and status code.
func (c *BaseGinController) HttpResponse(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}

// HttpReplySuccess sends user a http response with 200 status code and given data.
func (c *BaseGinController) HttpReplySuccess(context *gin.Context, data interface{}) {
	c.HttpResponse(context, gin.H{"data": data}, http.StatusOK)
}

// HttpReplyGinBindError sends user a http response with status code 422 and a set of validation errors.
// gennerally used when binding query params or json
func (c *BaseGinController) HttpReplyGinBindError(context *gin.Context, err error) {
	e := &controller.HttpError{
		Code:    422,
		Message: controller.ErrHttpInvalidRequest,
		Err:     nil,
		Errors:  err.Error(),
	}
	HandleHttpError(context, e)
}

// HttpReplyGinPathParamError sends user a http response with status code 400.
// used when any path parameter validation fails
func (c *BaseGinController) HttpReplyGinPathParamError(context *gin.Context, err error) {
	e := &controller.HttpError{
		Code:    400,
		Message: controller.ErrHttpInvalidPathParam,
		Err:     err,
		Errors:  nil,
	}
	HandleHttpError(context, e)
}

// HttpReplyGinPathParamError sends user a http response with status code 404.
func (c *BaseGinController) HttpReplyGinNotFoundError(context *gin.Context, err error) {
	e := &controller.HttpError{
		Code:    400,
		Message: controller.ErrHttpInvalidRequest,
		Err:     nil,
		Errors:  err.Error(),
	}
	HandleHttpError(context, e)
}

// HttpReplyHttpError bind controller.HttpError with http response
func (c *BaseGinController) HttpReplyHttpError(context *gin.Context, err *controller.HttpError) {
	HandleHttpError(context, err)
}

// HttpReplyServiceError converts ServiceError to controller.HttpError and sends http response
func (c *BaseGinController) HttpReplyServiceError(context *gin.Context, err *service.ServiceError) {
	e := controller.NewHttpErrorFromServiceError(err)
	HandleHttpError(context, e)
}

// NewBaseGinController creates a new BaseGinController
func NewBaseGinController(logger logger.GooglyLoggerInterface) *BaseGinController {
	con := controller.NewBaseController(logger)
	return &BaseGinController{
		con,
	}
}
