package dic

import (
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/account-management/app"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/rdb"
)

var Builder *di.Builder
var Container di.Container

const SentryClient = "sentry_client"
const Logger = "logger"
const Rdb = "rdb"

func InitContainer(apps []app.AppInterface) di.Container {
	builder := InitBuilder(apps)
	Container = builder.Build()
	return Container
}

func InitBuilder(apps []app.AppInterface) *di.Builder {
	Builder, _ = di.NewBuilder()
	RegisterServices(Builder)
	for _, app := range apps {
		app.Build(Builder)
	}
	return Builder
}

func RegisterServices(builder *di.Builder) {
	builder.Add(di.Def{
		Name: SentryClient,
		Build: func(ctn di.Container) (interface{}, error) {
			return logger.NewSentryClient(), nil
		},
	})

	builder.Add(di.Def{
		Name: Logger,
		Build: func(ctn di.Container) (interface{}, error) {
			return logger.NewLogger(), nil
		},
	})

	builder.Add(di.Def{
		Name: Rdb,
		Build: func(ctn di.Container) (interface{}, error) {
			return rdb.NewDb(), nil
		},
	})

}
