package todo

import (
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
