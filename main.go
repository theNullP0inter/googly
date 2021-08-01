package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/account-management/app"
	accounts_app "github.com/theNullP0inter/account-management/app_repository/accounts"
	"github.com/theNullP0inter/account-management/command"
	"github.com/theNullP0inter/account-management/dic"
)

var INSTALLED_APPS = []app.AppInterface{
	&accounts_app.AccountsApp{},
}

func main() {
	viper.AutomaticEnv()

	dic.InitContainer(INSTALLED_APPS)
	client := dic.Container.Get(dic.SentryClient).(*sentry.Client)
	if client != nil {
		func() {
			defer client.Recover(nil, nil, nil)
			command.Execute()
		}()
	} else {
		command.Execute()
	}
}
