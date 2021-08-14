package accounts

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"gorm.io/gorm"
)

type AccountResourceManagerInterface interface {
	resource.DbResourceManagerIntereface
}

type AccountResourceManager struct {
	*resource.RdbResourceManager
}

func NewAccountResourceManager(rdb *gorm.DB, logger logger.GooglyLoggerInterface) AccountResourceManagerInterface {
	var model Account
	queryBuilder := resource.NewPaginatedRdbListQueryBuilder(rdb, logger)
	rm := resource.NewRdbResourceManager(rdb, logger, model, queryBuilder).(*resource.RdbResourceManager)
	return &AccountResourceManager{rm}
}
