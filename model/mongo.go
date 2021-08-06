package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseMongoModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `json:"-" bson:"created_at"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at"`
}
