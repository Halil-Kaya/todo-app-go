package todo

import (
	"github.com/gofiber/fiber/v2"
	"todo/app/src/auth"
)

type TodoHttpHandler struct {
	authGuard auth.AuthGuard
}

func NewTodoHttpHandler(authGuard auth.AuthGuard) *TodoHttpHandler {
	return &TodoHttpHandler{
		authGuard,
	}
}

func (h *TodoHttpHandler) createTodo(ctx *fiber.Ctx) error {
	return nil
}

func (h *TodoHttpHandler) getTodos(ctx *fiber.Ctx) error {
	return nil
}

func (h *TodoHttpHandler) updateTodo(ctx *fiber.Ctx) error {
	return nil
}

func (h *TodoHttpHandler) deleteTodo(ctx *fiber.Ctx) error {
	return nil
}

func (h *TodoHttpHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/todo")
	appGroup.Get("/", h.authGuard.JWTGuard(h.getTodos))
	appGroup.Post("/", h.authGuard.JWTGuard(h.createTodo))
	appGroup.Patch("/:todoId", h.authGuard.JWTGuard(h.updateTodo))
	appGroup.Delete("/:todoId", h.authGuard.JWTGuard(h.deleteTodo))
}
