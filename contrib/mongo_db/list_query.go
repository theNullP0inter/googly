package mongo_db

import (
	"reflect"

	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/resource"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultPageSize = 30

// MongoListQueryBuilderInterface converts resource.ListQuery to filters needed for mongo
// ListQuery() should facilitate this
type MongoListQueryBuilder interface {
	ListQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions)
}

// PaginatedMongoListQueryBuilder is a pagination implementation for MongoListQueryBuilder
type PaginatedMongoListQueryBuilder interface {
	MongoListQueryBuilder
	PaginationQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions)
}

// BasePaginatedMongoListQueryBuilder is a base implementation for PaginatedMongoListQueryBuilder
type BasePaginatedMongoListQueryBuilder struct {
	Logger logger.GooglyLoggerInterface
}

// ListQuery should be implemented for MongoListQueryBuilder.
func (qb *BasePaginatedMongoListQueryBuilder) ListQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions) {
	return qb.PaginationQuery(parameters)
}

// PaginationQuery Converts resource.ListQuery to pagination filters needed for mongo
func (qb *BasePaginatedMongoListQueryBuilder) PaginationQuery(parameters resource.ListQuery) (bson.M, *options.FindOptions) {
	query := bson.M{}
	queryOptions := options.Find()

	// if params do not match the required format, log and return empty filter
	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		qb.Logger.Error("Unexpected type of parameters for PaginationQuery")
		return query, queryOptions
	}
	paginationParameters := val.FieldByName("PaginatedListQuery")
	hasPaginationParams := paginationParameters.IsValid() && !paginationParameters.IsNil()

	// Parsing page number
	var page int64
	page = 0
	if hasPaginationParams {
		pageValue := val.FieldByName("Page")
		if !pageValue.IsValid() || pageValue.Kind() != reflect.Int {
			qb.Logger.Info("Page in in invalid format, Using default value")
		} else {
			page = pageValue.Int()
		}
	}

	// Parsing page size
	var pageSize int64
	pageSize = DefaultPageSize
	if hasPaginationParams {
		pageSizeValue := val.FieldByName("PageSize")
		if !pageSizeValue.IsValid() || pageSizeValue.Kind() != reflect.Int {
			qb.Logger.Info("PageSize in in invalid format, Using default value")
		} else {
			pageSize = pageSizeValue.Int()
		}
	}

	// Applying Pagination to query
	limit := pageSize
	offset := page * pageSize

	queryOptions.SetSkip(offset)
	queryOptions.SetLimit(limit)

	// Parsing orderBy
	var orderBy = ""
	if hasPaginationParams {
		orderByValue := val.FieldByName("OrderBy")
		if orderByValue.IsValid() && orderByValue.Kind() == reflect.String {
			orderBy = orderByValue.String()
		}
	}

	// Return if order by is not present
	if orderBy == "" {
		return query, queryOptions
	}

	// Parsing orderBy descending or ascending
	var orderDesc = false
	if hasPaginationParams {
		orderDescValue := val.FieldByName("OrderDesc")
		if !orderDescValue.IsValid() || orderDescValue.Kind() != reflect.Bool {
			qb.Logger.Info("OrderDesc in in invalid format, Using default value")
		} else {
			orderDesc = orderDescValue.Bool()
		}
	}

	// Adding sort to query
	if len(orderBy) > 0 {
		if orderDesc {
			queryOptions.SetSort(bson.M{orderBy: -1})
		} else {
			queryOptions.SetSort(bson.M{orderBy: 1})
		}
	}

	return query, queryOptions
}

// NewBasePaginatedMongoListQueryBuilder creates a new BasePaginatedMongoListQueryBuilder
func NewBasePaginatedMongoListQueryBuilder(logger logger.GooglyLoggerInterface) *BasePaginatedMongoListQueryBuilder {
	return &BasePaginatedMongoListQueryBuilder{Logger: logger}
}
