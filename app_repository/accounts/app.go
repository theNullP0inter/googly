package accounts

import (
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/app"
	"github.com/theNullP0inter/account-management/dic"
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
				ctn.Get(dic.Rdb).(*gorm.DB),
				ctn.Get(dic.Logger).(*logrus.Logger),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: AccountsServiceName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountService(
				ctn.Get(dic.Logger).(*logrus.Logger),
				ctn.Get(AccountsResourceManagerName).(AccountResourceManagerInterface),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: AccountsControllerName,
		Build: func(ctn di.Container) (interface{}, error) {
			return NewAccountController(
				ctn.Get(AccountsServiceName).(AccountServiceInterface),
				ctn.Get(dic.Logger).(*logrus.Logger),
			), nil
		},
	})
}
