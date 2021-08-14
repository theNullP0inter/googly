package resource

import (
	"errors"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/db/model"
	"github.com/theNullP0inter/googly/logger"
	"gorm.io/gorm"
)

type RdbResourceManager struct {
	*ResourceManager
	Db           *gorm.DB
	Model        model.BaseModelInterface
	QueryBuilder RdbListQueryBuilderInterface
}

func handleGormError(err error) error {

	if err == gorm.ErrRecordNotFound {
		return ErrResourceNotFound
	} else if err == gorm.ErrInvalidTransaction {
		return ErrInvalidTransaction
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return ErrUniqueConstraint
	}
	return ErrInvalidQuery
}

func (s RdbResourceManager) Create(m DataInterface) (DataInterface, error) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	copier.Copy(item, m)
	result := s.Db.Create(item)

	if result.Error != nil {
		return nil, handleGormError(result.Error)
	}
	return item, nil
}

func (s RdbResourceManager) GetResource() Resource {
	return s.Model
}

func (s RdbResourceManager) GetModel() DataInterface {
	return s.Model
}
func (s RdbResourceManager) Get(id DataInterface) (DataInterface, error) {
	strId := id.(string)
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	binId, err := model.StringToBinID(strId)
	if err != nil {
		return nil, ErrInvalidFormat
	}
	bId, _ := binId.MarshalBinary()
	err = s.Db.Where("id = ?", bId).First(item).Error
	if err != nil {
		return nil, handleGormError(err)
	}
	return item, nil
}

func (s RdbResourceManager) Update(id DataInterface, data DataInterface) error {

	m := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	strId := id.(string)
	binId, err := model.StringToBinID(strId)
	if err != nil {
		return ErrInvalidFormat
	}
	bId, _ := binId.MarshalBinary()

	copier.Copy(m, data)

	result := s.Db.Model(s.GetModel()).Where("id = ?", bId).Updates(m)

	if result.Error != nil {
		return handleGormError(result.Error)
	}

	return nil
}
func (s RdbResourceManager) Delete(id DataInterface) error {
	item, err := s.Get(id)
	if err != nil {
		return err
	}
	s.Db.Delete(item)
	return nil
}

func (s RdbResourceManager) List(parameters DataInterface) (DataInterface, error) {

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()
	result, err := s.QueryBuilder.ListQuery(parameters)
	if err != nil {
		return nil, ErrInternal
	}
	result = result.Find(items)
	if result.Error != nil {
		return nil, handleGormError(result.Error)
	}

	return items, nil
}

func NewRdbResourceManager(
	db *gorm.DB,
	logger logger.GooglyLoggerInterface,
	model model.BaseModelInterface,
	queryBuilder PaginatedRdbListQueryBuilderInterface,
) DbResourceManagerIntereface {
	resourceManager := NewResourceManager(logger, model)
	return &RdbResourceManager{
		ResourceManager: resourceManager.(*ResourceManager),
		Db:              db,
		Model:           model,
		QueryBuilder:    queryBuilder,
	}

}
