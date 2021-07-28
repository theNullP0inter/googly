package dic

import (
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/boilerplate-go/logger"
	"github.com/theNullP0inter/boilerplate-go/rdb"
)

var Builder *di.Builder
var Container di.Container

const SentryClient = "sentry_client"
const Logger = "logger"
const Rdb = "rdb"

func InitContainer() di.Container {
	builder := InitBuilder()
	Container = builder.Build()
	return Container
}

func InitBuilder() *di.Builder {
	Builder, _ = di.NewBuilder()
	RegisterServices(Builder)
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
