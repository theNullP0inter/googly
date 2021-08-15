package mongo_db

import (
	"context"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/db"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

// MongoResourceManager is a DbResourcemanager connected to mongo_db
type MongoResourceManager interface {
	resource.DbResourceManager
}

// BaseMongoResourceManager is a base impelementation of MongoResourceManager
type BaseMongoResourceManager struct {
	*resource.BaseResourceManager
	Db             *mongo.Database
	CollectionName string
	Model          db.BaseModelInterface
	QueryBuilder   MongoListQueryBuilder
}

// GetResource will get you the model resource
func (s *BaseMongoResourceManager) GetResource() resource.Resource {
	return s.Model
}

// GetModel will get you the model resource
func (s *BaseMongoResourceManager) GetModel() resource.Resource {
	return s.Model
}

// Create creates an entry in with given data
func (s *BaseMongoResourceManager) Create(m resource.DataInterface) (resource.DataInterface, error) {
	ctx, cancel := initContext()
	defer cancel()

	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	copier.Copy(item, m)

	res, err := s.Db.Collection(s.CollectionName).InsertOne(ctx, item)
	if err != nil {
		return nil, resource.ErrInternal
	}

	itemBit, err := bson.Marshal(item)
	if err != nil {
		return nil, ErrParseBson
	}
	itemMap := bson.M{}
	bson.Unmarshal(itemBit, itemMap)
	itemMap["_id"] = res.InsertedID
	itemMapBit, _ := bson.Marshal(itemMap)
	bson.Unmarshal(itemMapBit, item)
	return item, nil
}

// Get gets 1 item with given _id
func (s *BaseMongoResourceManager) Get(id resource.DataInterface) (resource.DataInterface, error) {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, resource.ErrInvalidFormat
	}

	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	err = s.Db.Collection(s.CollectionName).FindOne(ctx, bson.M{"_id": objectId}).Decode(item)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, resource.ErrResourceNotFound
		}
		return nil, resource.ErrInternal
	}

	if item == nil {
		return nil, resource.ErrResourceNotFound
	}

	return item, nil

}

// Update updates 1 item with given _id & given data/update_set
func (s *BaseMongoResourceManager) Update(id resource.DataInterface, data resource.DataInterface) error {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return resource.ErrInvalidFormat
	}

	req, err := bson.Marshal(data)

	if err != nil {
		return ErrParseBson
	}

	reqDoc := bson.M{}
	bson.Unmarshal(req, reqDoc)

	res, err := s.Db.Collection(s.CollectionName).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": reqDoc})
	if err != nil {
		return resource.ErrInternal
	}

	if res.ModifiedCount == 0 {
		return resource.ErrNoModification
	}

	return nil
}

// Delete will delete 1 item with given _id
func (s *BaseMongoResourceManager) Delete(id resource.DataInterface) error {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return resource.ErrInvalidFormat
	}

	res, err := s.Db.Collection(s.CollectionName).DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return resource.ErrInternal
	}

	if res.DeletedCount == 0 {
		return resource.ErrResourceNotFound
	}

	return nil

}

// List will get you a list of items
//
// it uses QueryBuilder.ListQuery() to filter the documents
func (s *BaseMongoResourceManager) List(parameters resource.DataInterface) (resource.DataInterface, error) {

	ctx, cancel := initContext()
	defer cancel()

	query, opts := s.QueryBuilder.ListQuery(parameters)

	cur, err := s.Db.Collection(s.CollectionName).Find(ctx, query, opts)
	if err != nil {
		return nil, resource.ErrInternal
	}

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()

	cur.All(ctx, items)

	return items, nil
}

// NewMongoResourceManager creates a new MongoResourceManager
func NewMongoResourceManager(
	mongoDb *mongo.Database,
	collectionName string,
	logger logger.GooglyLoggerInterface,
	model db.BaseModelInterface,
	queryBuilder MongoListQueryBuilder,
) *BaseMongoResourceManager {
	resourceManager := resource.NewBaseResourceManager(logger, model)
	return &BaseMongoResourceManager{
		BaseResourceManager: resourceManager,
		Db:                  mongoDb,
		CollectionName:      collectionName,
		Model:               model,
		QueryBuilder:        queryBuilder,
	}

}
