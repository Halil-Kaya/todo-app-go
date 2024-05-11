package todo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"todo/app/src/model"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(client *mongo.Client) TodoRepository {
	db := client.Database("app")
	collection := db.Collection("todos")
	return TodoRepository{collection}
}

func (repo *TodoRepository) CreateTodo(dto *model.Todo) (*model.Todo, error) {
	_, err := repo.collection.InsertOne(context.Background(), dto)
	if err != nil {
		return nil, err
	}
	return dto, nil
}
