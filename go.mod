module github.com/theNullP0inter/googly

go 1.16

require (
	github.com/fsnotify/fsnotify v1.5.0 // indirect
	github.com/sarulabs/di/v2 v2.4.2
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20210908233432-aa78b53d3365 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/theNullP0inter/googly/logger => ./logger

replace github.com/theNullP0inter/googly/route => ./route

replace github.com/theNullP0inter/googly/service => ./service

replace github.com/theNullP0inter/googly/resource => ./resource

replace github.com/theNullP0inter/googly/controller => ./controller

replace github.com/theNullP0inter/googly/ingress => ./ingress

replace github.com/theNullP0inter/googly/tests => ./tests
