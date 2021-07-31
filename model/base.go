package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID BinID `gorm:"primaryKey; default: (UUID_TO_BIN(UUID()))"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	b.ID = BinID(id)
	return err
}
