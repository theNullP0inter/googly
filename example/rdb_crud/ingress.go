package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/example/rdb_crud/consts"
)

type MainIngressInterface interface{}

type MainIngress struct{}

func (i *MainIngress) Connect(cnt di.Container) *gin.Engine {
	router := gin.New()

	// gin.SetMode(viper.GetString("GIN_MODE"))

	// router := gin.New()
	router.Use(gin.Recovery())

	// client := cnt.Get(consts.SentryClient).(*sentry.Client)

	// if client != nil {
	// 	router.Use(sentrygin.New(sentrygin.Options{
	// 		Repanic: true,
	// 	}))
	// }

	// Display Swagger documentation
	router.StaticFile("doc/swagger.json", "/doc/swagger.json")
	config := &ginSwagger.Config{
		URL: "/doc/swagger.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// swagger:route GET /ping common getPing
	//
	// Ping
	//
	// Get Ping and reply Pong
	//
	//     Responses:
	//       200:
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Adding accounts
	accounts_controller := cnt.Get(consts.AccountsControllerName).(controller.GinControllerConnectorInterface)
	accounts_router := router.Group("/account")
	accounts_controller.AddRoutes(accounts_router)

	return router
}

func NewMainIngress() *MainIngress {
	return &MainIngress{}
}
