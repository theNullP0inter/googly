package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theNullP0inter/account-management/logger"
)

type HttpControllerConnectorInterface interface {
	AddRoutes(*gin.RouterGroup)
}

type HttpControllerInterface interface {
	ControllerInterface
	HttpResponse(*gin.Context, interface{}, int)
	HttpReplySuccess(*gin.Context, interface{})
	HttpReplyError(*gin.Context, string, int)
}

type HttpController struct {
	*Controller
	HttpControllerInterface
}

func (c HttpController) HttpReplySuccess(context *gin.Context, data interface{}) {
	c.HttpResponse(context, gin.H{"data": data}, http.StatusOK)
}

func (c HttpController) HttpReplyError(context *gin.Context, message string, code int) {
	c.HttpResponse(context, gin.H{"message": message}, code)
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
