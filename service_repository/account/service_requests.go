package account

import (
	"github.com/theNullP0inter/account-management/service"
	"github.com/thedevsaddam/govalidator"
)

type AccountServiceCreateRequest struct {
	Username string `json:"username"`
}

func (r *AccountServiceCreateRequest) Validate() *service.ServiceError {
	rules := govalidator.MapData{
		"username": []string{"required"},
	}

	opts := govalidator.Options{
		Data:  r,
		Rules: rules,
	}
	v := govalidator.New(opts)
	e := v.ValidateStruct()

	if len(e) > 0 {
		return service.NewServiceRequestValidationError(service.ServiceErrorMap(e))

	}
	return nil
}
