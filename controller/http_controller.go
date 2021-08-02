package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
)

type HttpControllerConnectorInterface interface {
	AddRoutes(*gin.RouterGroup)
}

type HttpControllerInterface interface {
	ControllerInterface
	HttpResponse(*gin.Context, interface{}, int)
	HttpReplySuccess(*gin.Context, interface{})
	HttpReplyError(*gin.Context, *errors.GooglyError)
}

type HttpController struct {
	*Controller
	HttpControllerInterface
}

func (c HttpController) HttpReplySuccess(context *gin.Context, data interface{}) {
	c.HttpResponse(context, gin.H{"data": data}, http.StatusOK)
}

func (c HttpController) HttpReplyInternalError(context *gin.Context, err interface{}) {
	if viper.GetString("GIN_MODE") == "release" {
		c.HttpResponse(context, gin.H{"error": "Not Found"}, 500)
		return
	}
	c.HttpResponse(context, gin.H{"error": err}, 500)
}

func (c HttpController) HttpReplyNotFound(context *gin.Context) {
	c.HttpResponse(context, gin.H{"error": "Not Found"}, 404)
}

func (c HttpController) HttpReplyValidationError(context *gin.Context, err interface{}) {
	c.HttpResponse(context, gin.H{"error": err}, 422)
}

func (c HttpController) HttpReplyErrorMessage(context *gin.Context, err *errors.GooglyError) {
	c.HttpResponse(context, gin.H{"error_message": err.Message}, err.Status)
}

func (c HttpController) HttpReplyGenericBadRequest(context *gin.Context) {
	c.HttpResponse(context, gin.H{"error": "Bad Request"}, 400)
}

func (c HttpController) HttpReplyBadRequestFromError(context *gin.Context, err error) {
	c.HttpResponse(context, gin.H{"error": err.Error()}, 400)
}

func (c HttpController) HttpReplyBindError(context *gin.Context, err error) {
	c.HttpReplyBadRequestFromError(context, err)
}

func (c HttpController) HttpResponse(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}

func NewHttpController(logger logger.LoggerInterface) HttpControllerInterface {
	controller := NewController(logger)
	return &HttpController{
		Controller: controller.(*Controller),
	}
}
