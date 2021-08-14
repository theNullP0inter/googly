package resource

import (
	"context"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/db/model"
	"github.com/theNullP0inter/googly/logger"
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

	reqDoc := bson.M{}
	bson.Unmarshal(req, reqDoc)

	res, err := s.Db.Collection(s.CollectionName).UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": reqDoc})
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
	mongoDb *mongo.Database,
	collectionName string,
	logger logger.GooglyLoggerInterface,
	model model.BaseModelInterface,
	queryBuilder MongoListQueryBuilderInterface,
) DbResourceManagerIntereface {
	resourceManager := NewResourceManager(logger, model)
	return &MongoResourceManager{
		ResourceManager: resourceManager.(*ResourceManager),
		Db:              mongoDb,
		CollectionName:  collectionName,
		Model:           model,
		QueryBuilder:    queryBuilder,
	}

}
