package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/example/rdb_crud/consts"
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
	accounts_controller := cnt.Get(consts.AccountsControllerName).(controller.GinControllerConnectorInterface)
	accounts_router := router.Group("/account")
	accounts_controller.AddRoutes(accounts_router)

	return router
}

func NewMainGinIngress(cnt di.Container, port int) *ingress.GinIngress {
	return ingress.NewGinIngress("serve_http", cnt, port, &MainGinConnector{})
}
