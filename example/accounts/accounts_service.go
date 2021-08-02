package main

import (
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/service"
)

type AccountServiceInterface interface {
	service.RdbCrudImplementorInterface
}

type AccountService struct {
	*service.CrudService
}

func NewAccountService(logger logger.LoggerInterface, rm AccountResourceManagerInterface) *AccountService {
	rdb_crud_implementor := service.NewRdbCrudImplementor(logger, rm)
	crud_service := service.NewCrudService(logger, rm, rdb_crud_implementor)
	return &AccountService{
		crud_service,
	}
}
