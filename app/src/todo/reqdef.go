package todo

import "time"

type TodoCreateDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TodoCreateAck struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoUpdateDto struct {
}

type TodoUpdateAck struct {
}

type TodoGetDto struct {
}

type TodoGetAck struct {
}

type TodoDeleteDto struct {
}

type TodoDeleteAck struct {
}
