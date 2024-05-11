package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"fullName"`
	Content   string             `bson:"fullName"`
	CreatedAt time.Time          `bson:"createdAt"`
}
