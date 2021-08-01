package resource_repository

import (
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/model"
	"github.com/theNullP0inter/account-management/resource_manager"
	"gorm.io/gorm"
)

type AccountResourceManagerInterface interface {
	resource_manager.ModelResourceManagerInterface
}

type AccountResourceManager struct {
	*resource_manager.ModelResourceManager
}

func NewAccountResourceManager(rdb *gorm.DB, logger *logrus.Logger) AccountResourceManagerInterface {
	var model model.Account
	query_builder := resource_manager.NewBaseListQueryBuilder(rdb, logger)
	rm := resource_manager.NewModelResourceManager(rdb, logger, model, query_builder)
	return &AccountResourceManager{rm}
}
