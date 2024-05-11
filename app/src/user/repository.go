package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
	"todo/app/src/model"
)

type UserRepository struct {
	collection *mongo.Collection
	logger     *zap.SugaredLogger
}

func NewUserRepository(client *mongo.Client, logger *zap.SugaredLogger) (*UserRepository, error) {
	db := client.Database("app")
	collection := db.Collection("users")
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"nickname": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}
	return &UserRepository{collection, logger}, nil
}

func (repo *UserRepository) CreateUser(dto UserCreateDto) (*model.User, error) {
	user := &model.User{
		Id:        primitive.NewObjectID(),
		Nickname:  dto.Nickname,
		FullName:  dto.FullName,
		Password:  dto.Password,
		CreatedAt: time.Now(),
	}

	_, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		repo.logger.Error("Error creating user:", err)
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) FindByNickname(nickname string) *model.User {
	var user *model.User
	repo.collection.FindOne(context.Background(), bson.M{"nickname": nickname}).Decode(&user)
	return user
}

func (repo *UserRepository) FindById(userId string) *model.User {
	var user *model.User
	objectId, _ := primitive.ObjectIDFromHex(userId)
	repo.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&user)
	return user
}
