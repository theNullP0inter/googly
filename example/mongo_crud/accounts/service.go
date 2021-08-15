package accounts

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

type AccountServiceInterface interface {
	service.CrudDbService
}

type AccountService struct {
	*service.BaseCrudDbService
}

func NewAccountService(logger logger.GooglyLoggerInterface, rm AccountResourceManagerInterface) AccountServiceInterface {
	ser := service.NewBaseCrudDbService(logger, rm)
	return &AccountService{
		ser,
	}
}
