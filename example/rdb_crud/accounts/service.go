package accounts

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

type AccountServiceInterface interface {
	service.DbCrudServiceInterface
}

type AccountService struct {
	*service.DbCrudService
}

func NewAccountService(logger logger.GooglyLoggerInterface, rm AccountResourceManagerInterface) AccountServiceInterface {
	crudService := service.NewDbCrudService(logger, rm)
	return &AccountService{
		crudService,
	}
}
