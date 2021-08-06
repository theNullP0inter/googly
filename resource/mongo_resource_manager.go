package resource

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	googly_errors "github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}

type MongoResourceManager struct {
	*ResourceManager
	Db             *mongo.Database
	CollectionName string
	Model          model.BaseModelInterface

	// TODO: Add New Query Builder
	QueryBuilder RdbListQueryBuilderInterface
}

func (s MongoResourceManager) GetResource() Resource {
	return s.Model
}

func (s MongoResourceManager) GetModel() DataInterface {
	return s.Model
}

func (s MongoResourceManager) Create(m DataInterface) (DataInterface, *googly_errors.GooglyError) {
	ctx, cancel := initContext()
	defer cancel()

	// item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	// copier.Copy(item, m)

	res, err := s.Db.Collection(s.CollectionName).InsertOne(ctx, &m)
	if err != nil {
		return nil, googly_errors.NewInternalError(err)
	}

	if item, ok := m.(*model.BaseMongoModel); ok {
		item.ID = res.InsertedID.(primitive.ObjectID)
		return item, nil
	}
	return m, nil

}

func (s MongoResourceManager) Get(id DataInterface) (DataInterface, *googly_errors.GooglyError) {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, googly_errors.NewObjectIdAssertionError(id)
	}

	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	s.Db.Collection(s.CollectionName).FindOne(ctx, bson.M{"_id": objectId}).Decode(&item)

	if item == nil {
		return nil, googly_errors.NewResourceNotFoundError("", nil)
	}

	return item, nil

}

func (s MongoResourceManager) Update(id DataInterface, data DataInterface) *googly_errors.GooglyError {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return googly_errors.NewObjectIdAssertionError(id)
	}

	req := data.(map[string]interface{})
	res, err := s.Db.Collection(s.CollectionName).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M(req)})
	if err != nil {
		return googly_errors.NewInternalError(err)
	}

	if res.ModifiedCount == 0 {
		return googly_errors.NewResourceNotFoundError(s.CollectionName, fmt.Errorf("%d Modified", res.ModifiedCount))
	}

	return nil
}

func (s MongoResourceManager) Delete(id DataInterface) *googly_errors.GooglyError {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return googly_errors.NewObjectIdAssertionError(id)
	}

	res, err := s.Db.Collection(s.CollectionName).DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return googly_errors.NewInternalError(err)
	}

	if res.DeletedCount == 0 {
		return googly_errors.NewResourceNotFoundError(s.CollectionName, errors.New("0 items Deleted"))
	}

	return nil

}

func (s MongoResourceManager) List(parameters DataInterface) (DataInterface, *googly_errors.GooglyError) {

	ctx, cancel := initContext()
	defer cancel()

	// TODO: Add Query from parameters
	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()
	cur, err := s.Db.Collection(s.CollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, googly_errors.NewInternalError(err)
	}
	err = cur.All(ctx, &items)
	if err != nil {
		return nil, googly_errors.NewInternalError(err)
	}
	return items, nil
}

func NewMongoResourceManager(
	mongo_db *mongo.Database,
	collection_name string,
	logger logger.LoggerInterface,
	model model.BaseModelInterface,

	query_builder PaginatedRdbListQueryBuilderInterface,
) DbResourceManagerIntereface {
	resource_manager := NewResourceManager(logger, model)
	return &MongoResourceManager{
		ResourceManager: resource_manager.(*ResourceManager),
		Db:              mongo_db,
		CollectionName:  collection_name,
		Model:           model,
		QueryBuilder:    query_builder,
	}

}
