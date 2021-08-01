package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModelInterface interface {
}

type BaseModel struct {
	ID        BinID          `gorm:"primaryKey; default: (UUID_TO_BIN(UUID()))" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	b.ID = BinID(id)
	return err
}
