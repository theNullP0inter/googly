package accounts

import (
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/googly"

	"github.com/theNullP0inter/googly/example/rdb_crud/consts"
	"github.com/theNullP0inter/googly/logger"
	"gorm.io/gorm"
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
				ctn.Get(consts.Rdb).(*gorm.DB),
				ctn.Get(consts.Logger).(logger.LoggerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: consts.AccountsServiceName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountService(
				ctn.Get(consts.Logger).(logger.LoggerInterface),
				ctn.Get(consts.AccountsResourceManagerName).(AccountResourceManagerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: consts.AccountsControllerName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountController(
				ctn.Get(consts.AccountsServiceName).(AccountServiceInterface),
				ctn.Get(consts.Logger).(logger.LoggerInterface),
			), nil
		},
	})
}
