package accounts

import (
	"github.com/theNullP0inter/googly/contrib/mongo_db"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountResourceManagerInterface interface {
	resource.DbResourceManagerIntereface
}

type AccountResourceManager struct {
	*mongo_db.MongoResourceManager
}

func NewAccountResourceManager(db *mongo.Database, logger logger.GooglyLoggerInterface) AccountResourceManagerInterface {
	var model Account
	listQueryBuilder := mongo_db.NewPaginatedMongoListQueryBuilder(logger)
	rm := mongo_db.NewMongoResourceManager(db, "accounts", logger, model, listQueryBuilder)
	return &AccountResourceManager{rm}
}
