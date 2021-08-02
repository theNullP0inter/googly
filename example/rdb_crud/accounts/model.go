package accounts

import "github.com/theNullP0inter/account-management/model"

type Account struct {
	model.BaseModel
	Username string `gorm:"unique" json:"username"`
}
