package service

import "github.com/gin-gonic/gin"

type HttpServiceInterface interface {
	AddRoutes(*gin.RouterGroup)
}
