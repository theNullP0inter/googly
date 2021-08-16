package accounts

import "github.com/theNullP0inter/googly/contrib/rdb"

type Account struct {
	rdb.RdbSoftDeleteBaseModel
	Username string `gorm:"unique" json:"username"`
}
