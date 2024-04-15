package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"todo/app/model"
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
		Nickname: dto.Nickname,
		FullName: dto.FullName,
		Password: dto.Password,
	}

	_, err := repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		repo.logger.Error("Error creating user:", err)
		return nil, err
	}
	return user, nil
}
