package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
)

type GinControllerIngressInterface interface {
	AddRoutes(*gin.RouterGroup)
}

type GinControllerInterface interface {
	ControllerInterface
	HttpResponse(*gin.Context, interface{}, int)
	HttpReplySuccess(*gin.Context, interface{})
	HttpReplyError(*gin.Context, *errors.GooglyError)
}

type GinController struct {
	*Controller
}

func (c GinController) HttpReplySuccess(context *gin.Context, data interface{}) {
	c.HttpResponse(context, gin.H{"data": data}, http.StatusOK)
}

func (c GinController) HttpReplyInternalError(context *gin.Context, err interface{}) {
	if viper.GetString("GIN_MODE") == "release" {
		c.HttpResponse(context, gin.H{"error": "Internal Server Error"}, 500)
		return
	}
	c.HttpResponse(context, gin.H{"error": err}, 500)
}

func (c GinController) HttpReplyNotFound(context *gin.Context) {
	c.HttpResponse(context, gin.H{"error": "Not Found"}, 404)
}

func (c GinController) HttpReplyValidationError(context *gin.Context, err interface{}) {
	c.HttpResponse(context, gin.H{"error": err}, 422)
}

func (c GinController) HttpReplyErrorMessage(context *gin.Context, err *errors.GooglyError) {
	c.HttpResponse(context, gin.H{"error_message": err.Message}, err.Status)
}

func (c GinController) HttpReplyGenericBadRequest(context *gin.Context) {
	c.HttpResponse(context, gin.H{"error": "Bad Request"}, 400)
}

func (c GinController) HttpReplyBadRequestFromError(context *gin.Context, err error) {
	c.HttpResponse(context, gin.H{"error": err.Error()}, 400)
}

func (c GinController) HttpReplyBindError(context *gin.Context, err error) {
	c.HttpReplyBadRequestFromError(context, err)
}

func (c GinController) HttpResponse(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}

func NewGinController(logger logger.LoggerInterface) *GinController {
	controller := NewController(logger)
	return &GinController{
		controller,
	}
}
