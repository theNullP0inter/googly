package app_repository

import (
	"github.com/gin-gonic/gin"
	accounts_app "github.com/theNullP0inter/account-management/app_repository/accounts"
	"github.com/theNullP0inter/account-management/controller"
	"github.com/theNullP0inter/account-management/dic"
)

func RegisterRoutes(router *gin.RouterGroup) {

	// Adding accounts
	accounts_controller := dic.Container.Get(accounts_app.AccountsControllerName).(controller.HttpControllerConnectorInterface)
	accounts_router := router.Group("/account")
	accounts_controller.AddRoutes(accounts_router)

}
