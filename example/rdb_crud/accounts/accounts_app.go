package accounts

import (
	"github.com/sarulabs/di/v2"
	"github.com/theNullP0inter/account-management/app"

	"github.com/theNullP0inter/account-management/example/rdb_crud"
	"github.com/theNullP0inter/account-management/logger"
	"gorm.io/gorm"
)

type AccountsAppInterface interface {
	app.AppInterface
}

type AccountsApp struct {
	AccountsAppInterface
	ResourceManager *AccountResourceManager
	Service         *AccountService
	Controller      *AccountController
}

const AccountsServiceName = "accounts_service"
const AccountsResourceManagerName = "accounts_resource_management"
const AccountsControllerName = "accounts_controller"

func (a *AccountsApp) Build(builder *di.Builder) {

	builder.Add(di.Def{
		Name: AccountsResourceManagerName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountResourceManager(
				ctn.Get(rdb_crud.Rdb).(*gorm.DB),
				ctn.Get(rdb_crud.Logger).(logger.LoggerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: AccountsServiceName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountService(
				ctn.Get(rdb_crud.Logger).(logger.LoggerInterface),
				ctn.Get(AccountsResourceManagerName).(AccountResourceManagerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: AccountsControllerName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountController(
				ctn.Get(AccountsServiceName).(AccountServiceInterface),
				ctn.Get(rdb_crud.Logger).(logger.LoggerInterface),
			), nil
		},
	})
}
