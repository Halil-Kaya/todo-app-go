package exception

import "fmt"

type TodoNotFound struct {
	Code    int
	Message string
}

func NewTodoNotFound() *TodoNotFound {
	return &TodoNotFound{Code: 404001, Message: "Todo Not Found"}
}

func (e *TodoNotFound) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *TodoNotFound) GetCode() int {
	return e.Code
}

func (e *TodoNotFound) GetMessage() string {
	return e.Message
}
