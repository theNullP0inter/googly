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

func NewAccountResourceManager(db *mongo.Database, logger logger.LoggerInterface) AccountResourceManagerInterface {
	var model Account
	list_query_builder := resource.NewPaginatedMongoListQueryBuilder(logger)
	rm := resource.NewMongoResourceManager(db, "accounts", logger, model, list_query_builder).(*resource.MongoResourceManager)
	return &AccountResourceManager{rm}
}
