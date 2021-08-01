package dic

import (
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/controller_repository"
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/rdb"
	resource_repository "github.com/theNullP0inter/account-management/resource_repository"
	service_repository "github.com/theNullP0inter/account-management/service_repository"
	"gorm.io/gorm"
)

var Builder *di.Builder
var Container di.Container

const SentryClient = "sentry_client"
const Logger = "logger"
const Rdb = "rdb"

const AccountService = "account_service"
const AccountResourceManager = "account_resource_manager"
const AccountController = "account_controller"

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
		Name: AccountResourceManager,
		Build: func(ctn di.Container) (interface{}, error) {
			return resource_repository.NewAccountResourceManager(
				ctn.Get(Rdb).(*gorm.DB),
				ctn.Get(Logger).(*logrus.Logger),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: AccountService,
		Build: func(ctn di.Container) (interface{}, error) {
			return service_repository.NewAccountService(
				ctn.Get(Logger).(*logrus.Logger),
				ctn.Get(AccountResourceManager).(resource_repository.AccountResourceManagerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: AccountController,
		Build: func(ctn di.Container) (interface{}, error) {
			return controller_repository.NewAccountController(
				ctn.Get(AccountService).(service_repository.AccountServiceInterface),
				ctn.Get(Logger).(*logrus.Logger),
			), nil
		},
	})

}
