package accounts

import (
	"github.com/theNullP0inter/googly/contrib/mongo_db"
	"github.com/theNullP0inter/googly/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountResourceManagerInterface interface {
	mongo_db.MongoResourceManager
}

type AccountResourceManager struct {
	*mongo_db.BaseMongoResourceManager
}

func NewAccountResourceManager(db *mongo.Database, logger logger.GooglyLoggerInterface) AccountResourceManagerInterface {
	var model Account
	listQueryBuilder := mongo_db.NewBasePaginatedMongoListQueryBuilder(logger)
	rm := mongo_db.NewMongoResourceManager(db, "accounts", logger, model, listQueryBuilder)
	return &AccountResourceManager{rm}
}
