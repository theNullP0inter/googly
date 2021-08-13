package accounts

import (
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/googly"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/theNullP0inter/googly/example/mongo_crud/consts"
	"github.com/theNullP0inter/googly/logger"
)

type AccountsAppInterface interface {
	googly.AppInterface
}

type AccountsApp struct {
	AccountsAppInterface
	ResourceManager *AccountResourceManager
	Service         *AccountService
	Controller      *AccountController
}

func (a *AccountsApp) Build(builder *di.Builder) {

	builder.Add(di.Def{
		Name: consts.AccountsResourceManagerName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountResourceManager(
				ctn.Get(consts.MongoDb).(*mongo.Database),
				ctn.Get(consts.Logger).(logger.GooglyLoggerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: consts.AccountsServiceName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountService(
				ctn.Get(consts.Logger).(logger.GooglyLoggerInterface),
				ctn.Get(consts.AccountsResourceManagerName).(AccountResourceManagerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: consts.AccountsControllerName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountController(
				ctn.Get(consts.AccountsServiceName).(AccountServiceInterface),
				ctn.Get(consts.Logger).(logger.GooglyLoggerInterface),
			), nil
		},
	})
}
