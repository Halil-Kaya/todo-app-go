package todo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repo *TodoRepository) GetTodos(userId primitive.ObjectID) ([]*model.Todo, error) {
	cursor, err := repo.collection.Find(context.Background(), bson.M{"userid": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var todos []*model.Todo
	for cursor.Next(context.Background()) {
		var todo model.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
