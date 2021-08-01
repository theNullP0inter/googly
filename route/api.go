// App worker
//
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
// swagger:meta
package route

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/theNullP0inter/account-management/dic"
)

func Setup(builder *di.Builder) *gin.Engine {
	gin.SetMode(viper.GetString("GIN_MODE"))

	router := gin.New()
	router.Use(gin.Recovery())

	client := dic.Container.Get(dic.SentryClient).(*sentry.Client)

	if client != nil {
		router.Use(sentrygin.New(sentrygin.Options{
			Repanic: true,
		}))
	}

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
	return router
}
