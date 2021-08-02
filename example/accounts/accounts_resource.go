package main

import (
	"github.com/theNullP0inter/account-management/logger"
	"github.com/theNullP0inter/account-management/resource"
	"gorm.io/gorm"
)

type AccountResourceManagerInterface interface {
	resource.RdbCrudResourceManagerIntereface
}

type AccountResourceManager struct {
	*resource.RdbCrudResourceManager
}

func NewAccountResourceManager(rdb *gorm.DB, logger logger.LoggerInterface) AccountResourceManagerInterface {
	var model Account
	query_builder := resource.NewPaginatedRdbListQueryBuilder(rdb, logger)
	rm := resource.NewRdbCrudResourceManager(rdb, logger, model, query_builder)
	return &AccountResourceManager{rm}
}
