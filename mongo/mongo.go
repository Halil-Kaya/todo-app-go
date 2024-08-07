package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo/config"
	"todo/logger"
)

func ConnectDb(config config.Config) *mongo.Client {
	clientOptions := options.Client().ApplyURI(config.MongoDB.ConnectionUrl)
	client, err := mongo.Connect(context.Background(), clientOptions)
	logger := logger.GetLogger()
	if err != nil {
		logger.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("MongoDb connection is successful")

	return client
}
