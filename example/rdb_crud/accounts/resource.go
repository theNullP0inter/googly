package accounts

import (
	"github.com/theNullP0inter/googly/contrib/rdb"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"gorm.io/gorm"
)

type AccountResourceManagerInterface interface {
	resource.DbResourceManagerIntereface
}

type AccountResourceManager struct {
	*rdb.RdbResourceManager
}

func NewAccountResourceManager(db *gorm.DB, logger logger.GooglyLoggerInterface) AccountResourceManagerInterface {
	var model Account
	queryBuilder := rdb.NewPaginatedRdbListQueryBuilder(db, logger)
	rm := rdb.NewRdbResourceManager(db, logger, model, queryBuilder)
	return &AccountResourceManager{rm}
}
