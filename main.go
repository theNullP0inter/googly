package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"github.com/theNullP0inter/boilerplate-go/command"
	"github.com/theNullP0inter/boilerplate-go/dic"
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
