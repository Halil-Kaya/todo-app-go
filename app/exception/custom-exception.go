package exception

type ICustomException interface {
	GetCode() int
	GetMessage() string
	Error() string
}
