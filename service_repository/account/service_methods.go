package account

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/copier"
	"github.com/theNullP0inter/account-management/model"
	"github.com/theNullP0inter/account-management/service"
)

func (s *AccountService) Create(req *AccountServiceCreateRequest) (*AccountResource, *service.ServiceError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account_model := model.Account{
		Username: req.Username,
	}
	result := s.Rdb.Create(&account_model)
	if result.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, service.NewUniqueConstraintError("account")
		}
		return nil, service.NewInternalServiceError(result.Error)
	}

	account_resource := &AccountResource{}
	copier.Copy(&account_resource, &account_model)

	return account_resource, nil
}

func (s *AccountService) Query() (*[]AccountResource, *service.ServiceError) {

	var account_records []model.Account
	result := s.Rdb.Find(&account_records)
	if result.Error != nil {
		return nil, service.NewInternalServiceError(result.Error)
	}

	account_resources := []AccountResource{}
	copier.Copy(&account_resources, &account_records)

	return &account_resources, nil
}

func (s *AccountService) Delete(id model.BinID) *service.ServiceError {
	b_id, _ := id.MarshalBinary()
	result := s.Rdb.Where("id = ? ", b_id).Delete(&model.Account{})

	if result.Error != nil {
		return service.NewInternalServiceError(result.Error)
	}

	return nil
}
