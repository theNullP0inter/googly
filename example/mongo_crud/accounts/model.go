package accounts

import "github.com/theNullP0inter/googly/model"

type Account struct {
	model.BaseMongoModel
	Username string `bson:"username" json:"username"`
}
