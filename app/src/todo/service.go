package todo

import (
	"go.mongodb.org/mongo-driver/bson"
	"todo/app/core/exception"
	"todo/app/src/model"
)

type TodoService struct {
	todoRepository TodoRepository
}

func NewTodoService(todoRepository TodoRepository) TodoService {
	return TodoService{todoRepository}
}

func (service *TodoService) CreateTodo(dto *model.Todo) (*model.Todo, error) {
	createdTodo, err := service.todoRepository.CreateTodo(dto)
	if err != nil {
		return nil, err
	}
	return createdTodo, nil
}

func (service *TodoService) UpdateTodo(todoId string, dto TodoUpdateDto) error {
	todo := service.todoRepository.FindById(todoId)
	if todo == nil {
		return exception.NewTodoNotFound()
	}
	updateFields := bson.D{}
	if dto.Title != nil {
		updateFields = append(updateFields, bson.E{"title", *dto.Title})
	}
	if dto.Content != nil {
		updateFields = append(updateFields, bson.E{"content", *dto.Content})
	}

	err := service.todoRepository.UpdateTodo(todo.Id, updateFields)
	return err
}
