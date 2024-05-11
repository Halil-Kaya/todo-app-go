package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Nickname  string             `bson:"nickname" unique:"true"`
	FullName  string             `bson:"fullName"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"createdAt"`
}
