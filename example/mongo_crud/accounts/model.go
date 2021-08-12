package accounts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"-" `
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"-" `
	Username  string             `bson:"username,omitempty" json:"username"`
}
