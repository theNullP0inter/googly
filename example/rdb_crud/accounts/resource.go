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

func NewAccountResourceManager(rdb *gorm.DB, logger logger.LoggerInterface) AccountResourceManagerInterface {
	var model Account
	query_builder := resource.NewPaginatedRdbListQueryBuilder(rdb, logger)
	rm := resource.NewRdbResourceManager(rdb, logger, model, query_builder).(*resource.RdbResourceManager)
	return &AccountResourceManager{rm}
}
