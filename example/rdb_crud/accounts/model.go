package accounts

import "github.com/theNullP0inter/googly/model"

type Account struct {
	model.RdbSoftDeleteBaseModel
	Username string `gorm:"unique" json:"username"`
}
