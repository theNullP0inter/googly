package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HttpConnectorInterface interface {
	AddRoutes(*gin.RouterGroup)
}

type HttpController struct {
	*Controller
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

func NewHttpController(logger *logrus.Logger) *HttpController {
	controller := NewController(logger)
	return &HttpController{
		controller,
	}
}
