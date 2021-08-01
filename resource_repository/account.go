package resource_repository

import (
	"github.com/sirupsen/logrus"
	"github.com/theNullP0inter/account-management/model"
	"github.com/theNullP0inter/account-management/resource"
	"gorm.io/gorm"
)

type AccountResourceManagerInterface interface {
	resource.ModelResourceManagerInterface
}

type AccountResourceManager struct {
	*resource.ModelResourceManager
}

func NewAccountResourceManager(rdb *gorm.DB, logger *logrus.Logger) AccountResourceManagerInterface {
	var model model.Account
	query_builder := resource.NewBaseListQueryBuilder(rdb, logger)
	rm := resource.NewModelResourceManager(rdb, logger, model, query_builder)
	return &AccountResourceManager{rm}
}
