package mongo_db

import (
	"reflect"

	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultPageSize = 30

type MongoListQueryBuilderInterface interface {
	ListQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions)
}
type PaginatedMongoListQueryBuilderInterface interface {
	MongoListQueryBuilderInterface
	PaginationQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions)
}

type PaginatedMongoListQueryBuilder struct {
	Logger logger.GooglyLoggerInterface
	PaginatedMongoListQueryBuilderInterface
}

func NewPaginatedMongoListQueryBuilder(logger logger.GooglyLoggerInterface) *PaginatedMongoListQueryBuilder {
	return &PaginatedMongoListQueryBuilder{Logger: logger}
}

// modify for a new type of query builder
func (c PaginatedMongoListQueryBuilder) ListQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions) {
	return c.PaginationQuery(parameters)
}

func (c PaginatedMongoListQueryBuilder) PaginationQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions) {
	query := bson.M{}
	queryOptions := options.Find()

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		c.Logger.Errorf("Unexpected type of parameters for PaginationQuery")
		return query, queryOptions
	}
	paginationParameters := val.FieldByName("PaginatedListQuery")
	hasPaginationParams := paginationParameters.IsValid() && !paginationParameters.IsNil()

	var page int64
	page = 0
	if hasPaginationParams {
		pageValue := val.FieldByName("Page")
		if !pageValue.IsValid() || pageValue.Kind() != reflect.Int {
			c.Logger.Errorf("Page is not specified correctly in listQuery")
		} else {
			page = pageValue.Int()
		}
	}

	var pageSize int64
	pageSize = DefaultPageSize
	if hasPaginationParams {
		pageSizeValue := val.FieldByName("PageSize")
		if !pageSizeValue.IsValid() || pageSizeValue.Kind() != reflect.Int {
			c.Logger.Errorf("PageSize is not specified in listQuery")
		} else {
			pageSize = pageSizeValue.Int()
		}
	}

	limit := pageSize
	offset := page * pageSize

	queryOptions.SetSkip(offset)
	queryOptions.SetLimit(limit)

	var orderBy = "_id"
	if hasPaginationParams {
		orderByValue := val.FieldByName("OrderBy")
		if orderByValue.IsValid() && orderByValue.Kind() == reflect.String {
			orderBy = orderByValue.String()
		}
	}

	var orderDesc = false
	if hasPaginationParams {
		orderDescValue := val.FieldByName("OrderDesc")
		if orderDescValue.IsValid() && orderDescValue.Kind() == reflect.Bool {
			orderDesc = orderDescValue.Bool()
		}
	}

	if len(orderBy) > 0 {
		if orderDesc {
			queryOptions.SetSort(bson.M{orderBy: -1})
		} else {
			queryOptions.SetSort(bson.M{orderBy: 1})
		}
	}

	return query, queryOptions
}
