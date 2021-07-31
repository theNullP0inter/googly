package account

import "github.com/theNullP0inter/account-management/model"

type AccountResource struct {
	ID       model.BinID `copier:"must" json:"id"`
	Username string      `copier:"must" json:"username"`
}
