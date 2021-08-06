package resource

import (
	go_errors "errors"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/googly/errors"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/model"
	"gorm.io/gorm"
)

type RdbResourceManager struct {
	*ResourceManager
	Db           *gorm.DB
	Model        model.BaseModelInterface
	QueryBuilder RdbListQueryBuilderInterface
}

func handleGormError(err error) *errors.GooglyError {

	if err == gorm.ErrRecordNotFound {
		return errors.NewResourceNotFoundError("", err)
	} else if err == gorm.ErrInvalidTransaction {
		return errors.NewInternalError(err)
	}

	var mysqlErr *mysql.MySQLError
	if go_errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errors.NewUniqueConstraintError("", err)
	}
	return errors.NewInvalidRequestError()
}

func (s RdbResourceManager) Create(m DataInterface) (DataInterface, *errors.GooglyError) {
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
func (s RdbResourceManager) Get(id DataInterface) (DataInterface, *errors.GooglyError) {
	item := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	bin_id, ok := id.(model.BinID)
	if !ok {
		return nil, errors.NewBinIdAssertionError(id)
	}
	b_id, _ := bin_id.MarshalBinary()
	err := s.Db.Where("id = ?", b_id).First(item).Error
	if err != nil {
		return nil, handleGormError(err)
	}
	return item, nil
}

func (s RdbResourceManager) Update(id DataInterface, data DataInterface) *errors.GooglyError {

	m := reflect.New(reflect.TypeOf(s.GetModel())).Interface()
	bin_id, ok := id.(model.BinID)
	if !ok {
		return errors.NewBinIdAssertionError(id)
	}
	b_id, _ := bin_id.MarshalBinary()

	copier.Copy(m, data)

	result := s.Db.Model(s.GetModel()).Where("id = ?", b_id).Updates(m)

	if result.Error != nil {
		return handleGormError(result.Error)
	}

	return nil
}
func (s RdbResourceManager) Delete(id DataInterface) *errors.GooglyError {
	item, err := s.Get(id)
	if err != nil {
		return err
	}
	s.Db.Delete(item)
	return nil
}

func (s RdbResourceManager) List(parameters DataInterface) (DataInterface, *errors.GooglyError) {

	items := reflect.New(reflect.SliceOf(reflect.TypeOf(s.GetModel()))).Interface()
	result, err := s.QueryBuilder.ListQuery(parameters)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}
	result = result.Find(items)
	if result.Error != nil {
		return nil, handleGormError(result.Error)
	}

	return items, nil
}

func NewRdbResourceManager(
	db *gorm.DB,
	logger logger.LoggerInterface,
	model model.BaseModelInterface,
	query_builder PaginatedRdbListQueryBuilderInterface,
) DbResourceManagerIntereface {
	resource_manager := NewResourceManager(logger, model)
	return &RdbResourceManager{
		ResourceManager: resource_manager.(*ResourceManager),
		Db:              db,
		Model:           model,
		QueryBuilder:    query_builder,
	}

}
