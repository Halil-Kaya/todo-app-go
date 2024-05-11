package user

type UserCreateDto struct {
	Nickname string `json:"nickname" validate:"required,min=2"`
	FullName string `json:"fullname" validate:"required,min=2"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserCreateAck struct {
	Id       string
	Nickname string
}

type UserMeAck struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	FullName string `json:"fullName"`
}
