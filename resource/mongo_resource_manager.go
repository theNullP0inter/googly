package resource

import (
	"context"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
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
	QueryBuilder   MongoListQueryBuilderInterface
}

func (s MongoResourceManager) GetResource() Resource {
	return s.Model
}

func (s MongoResourceManager) GetModel() DataInterface {
	return s.Model
}

func (s MongoResourceManager) Create(m DataInterface) (DataInterface, error) {
	ctx, cancel := initContext()
	defer cancel()

	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	copier.Copy(item, m)

	res, err := s.Db.Collection(s.CollectionName).InsertOne(ctx, item)
	if err != nil {
		return nil, ErrInternal
	}

	item_bit, err := bson.Marshal(item)
	if err != nil {
		return nil, ErrParseBson
	}
	item_map := bson.M{}
	bson.Unmarshal(item_bit, item_map)
	item_map["_id"] = res.InsertedID
	item_map_bit, _ := bson.Marshal(item_map)
	bson.Unmarshal(item_map_bit, item)
	return item, nil
}

func (s MongoResourceManager) Get(id DataInterface) (DataInterface, error) {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, ErrInvalidFormat
	}

	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	err = s.Db.Collection(s.CollectionName).FindOne(ctx, bson.M{"_id": objectId}).Decode(item)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrResourceNotFound
		}
		return nil, ErrInternal
	}

	if item == nil {
		return nil, ErrResourceNotFound
	}

	return item, nil

}

func (s MongoResourceManager) Update(id DataInterface, data DataInterface) error {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return ErrInvalidFormat
	}

	req, err := bson.Marshal(data)

	if err != nil {
		return ErrParseBson
	}

	req_doc := bson.M{}
	bson.Unmarshal(req, req_doc)

	res, err := s.Db.Collection(s.CollectionName).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": req_doc})
	if err != nil {
		return ErrInternal
	}

	if res.ModifiedCount == 0 {
		return ErrNoModification
	}

	return nil
}

func (s MongoResourceManager) Delete(id DataInterface) error {
	ctx, cancel := initContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return ErrInvalidFormat
	}

	res, err := s.Db.Collection(s.CollectionName).DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		return ErrInternal
	}

	if res.DeletedCount == 0 {
		return ErrResourceNotFound
	}

	return nil

}

func (s MongoResourceManager) List(parameters DataInterface) (DataInterface, error) {

	ctx, cancel := initContext()
	defer cancel()

	query, opts := s.QueryBuilder.ListQuery(parameters)

	cur, err := s.Db.Collection(s.CollectionName).Find(ctx, query, opts)
	if err != nil {
		return nil, ErrInternal
	}

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()

	cur.All(ctx, items)

	return items, nil
}

func NewMongoResourceManager(
	mongo_db *mongo.Database,
	collection_name string,
	logger logger.GooglyLoggerInterface,
	model model.BaseModelInterface,
	query_builder MongoListQueryBuilderInterface,
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
