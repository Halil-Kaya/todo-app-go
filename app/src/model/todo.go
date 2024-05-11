package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	UserId    primitive.ObjectID `json:"userId"`
	CreatedAt time.Time          `bson:"createdAt"`
}
