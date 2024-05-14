package todo

import "time"

type TodoCreateDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoCreateAck struct {
	Id        string    `json:"id"`
	Title     string    `json:"title" validate:"required,min=2"`
	Content   string    `json:"content" validate:"required,min=2"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoUpdateDto struct {
	Title   *string `json:"title" validate:"min=2"`
	Content *string `json:"content" validate:"min=2"`
}

type TodoUpdateAck struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoGetDto struct {
}

type TodoAck struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoGetAck struct {
	Todos []TodoAck `json:"todos"`
}

type TodoDeleteDto struct {
}

type TodoDeleteAck struct {
}
