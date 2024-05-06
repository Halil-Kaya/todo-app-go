package exception

import "fmt"

type NicknameIsAlreadyTaken struct {
	ICustomException
	Code    int
	Message string
}

func NewNicknameIsAlreadyTaken() *NicknameIsAlreadyTaken {
	return &NicknameIsAlreadyTaken{Code: 500100, Message: "Nickname Is Already Taken"}
}

func (e *NicknameIsAlreadyTaken) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *NicknameIsAlreadyTaken) GetCode() int {
	return e.Code
}

func (e *NicknameIsAlreadyTaken) GetMessage() string {
	return e.Message
}
