package model

type Account struct {
	BaseModel
	Username string `gorm:"unique"`
}
