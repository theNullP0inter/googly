package dic

import (
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/rdb"
	account_service "github.com/theNullP0inter/account-management/service_repository/account"
	"gorm.io/gorm"
)

var Builder *di.Builder
var Container di.Container

const SentryClient = "sentry_client"
const Logger = "logger"
const Rdb = "rdb"

const AccountService = "account_service"

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

	builder.Add(di.Def{
		Name: AccountService,
		Build: func(ctn di.Container) (interface{}, error) {
			return account_service.NewAccountService(
				ctn.Get(Rdb).(*gorm.DB),
				ctn.Get(Logger).(*logrus.Logger),
			), nil
		},
	})

}
