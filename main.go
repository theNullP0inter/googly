package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/account-management/command"
	"github.com/theNullP0inter/account-management/dic"
)

func main() {
	viper.AutomaticEnv()

	dic.InitContainer()
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
