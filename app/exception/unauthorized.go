package exception

import "fmt"

type Unauthorized struct {
	ICustomException
	Code    int
	Message string
}

func NewUnauthorized() *Unauthorized {
	return &Unauthorized{Code: 401100, Message: "Unauthorized"}
}

func (e *Unauthorized) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *Unauthorized) GetCode() int {
	return e.Code
}

func (e *Unauthorized) GetMessage() string {
	return e.Message
}
