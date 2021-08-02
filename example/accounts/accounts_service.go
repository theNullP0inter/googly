package main

import (
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/service"
)

type AccountServiceInterface interface {
	service.DbCrudServiceInterface
}

type AccountService struct {
	*service.DbCrudService
}

func NewAccountService(logger logger.LoggerInterface, rm AccountResourceManagerInterface) AccountServiceInterface {
	crud_service := service.NewDbCrudService(logger, rm)
	return &AccountService{
		crud_service.(*service.DbCrudService),
	}
}
