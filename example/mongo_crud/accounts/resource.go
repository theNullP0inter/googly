package accounts

import (
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountResourceManagerInterface interface {
	resource.DbResourceManagerIntereface
}

type AccountResourceManager struct {
	*resource.MongoResourceManager
}

func NewAccountResourceManager(db *mongo.Database, logger logger.GooglyLoggerInterface) AccountResourceManagerInterface {
	var model Account
	listQueryBuilder := resource.NewPaginatedMongoListQueryBuilder(logger)
	rm := resource.NewMongoResourceManager(db, "accounts", logger, model, listQueryBuilder).(*resource.MongoResourceManager)
	return &AccountResourceManager{rm}
}
