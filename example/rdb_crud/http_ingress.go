package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	gin_contrib "github.com/theNullP0inter/googly/contrib/gin"
	"github.com/theNullP0inter/googly/example/rdb_crud/consts"
)

type MainGinConnector struct{}

func (i *MainGinConnector) Connect(cnt di.Container) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Adding accounts
	accountsController := cnt.Get(consts.AccountsControllerName).(gin_contrib.GinControllerIngress)
	accountsRouter := router.Group("/account")
	accountsController.AddRoutes(accountsRouter)

	return router
}

func NewMainGinIngress(cnt di.Container, port int) *gin_contrib.GinIngress {
	return gin_contrib.NewGinIngress("serve_http", cnt, port, &MainGinConnector{})
}
