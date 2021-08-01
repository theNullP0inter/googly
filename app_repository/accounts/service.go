package accounts

import (
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/resource"
	"github.com/theNullP0inter/account-management/service"
)

type AccountServiceInterface interface {
	service.ModelCrudServiceInterface
}

type AccountService struct {
	*service.ModelCrudService
}

func NewAccountService(logger *logrus.Logger, rm resource.ModelResourceManagerInterface) *AccountService {
	model_crud_service := service.NewModelCrudService(logger, rm)
	return &AccountService{
		model_crud_service,
	}
}
