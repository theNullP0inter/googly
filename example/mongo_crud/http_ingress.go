package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/example/mongo_crud/consts"
	"github.com/theNullP0inter/googly/ingress"
)

type MainGinConnector struct{}

func (i *MainGinConnector) Connect(cnt di.Container) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Adding accounts
	accountsController := cnt.Get(consts.AccountsControllerName).(controller.GinControllerIngress)
	accountsRouter := router.Group("/account")
	accountsController.AddRoutes(accountsRouter)

	return router
}

func NewMainGinIngress(cnt di.Container, port int) *ingress.GinIngress {
	return ingress.NewGinIngress("serve_http", cnt, port, &MainGinConnector{})
}
