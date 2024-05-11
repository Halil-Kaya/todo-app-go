package todo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todo/app/core/utility"
	"todo/app/core/validation"
	"todo/app/src/auth"
	"todo/app/src/model"
)

type TodoHttpHandler struct {
	authGuard   auth.AuthGuard
	todoService TodoService
}

func NewTodoHttpHandler(authGuard auth.AuthGuard, todoService TodoService) *TodoHttpHandler {
	return &TodoHttpHandler{
		authGuard,
		todoService,
	}
}

func (h *TodoHttpHandler) createTodo(ctx *fiber.Ctx) error {
	var request TodoCreateDto
	err := ctx.BodyParser(&request)
	if err != nil {
		return utility.ErrorResponse(ctx, err)
	}
	if errors := validation.Validate(request); len(errors) > 0 {
		return utility.ErrorResponse(ctx, errors)

	}
	user := ctx.Locals("user").(*model.User)

	model := &model.Todo{
		Id:        primitive.NewObjectID(),
		Title:     request.Title,
		Content:   request.Content,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}

	createdTodo, err := h.todoService.CreateTodo(model)

	if err != nil {
		return utility.ErrorResponse(ctx, err)
	}

	return utility.OkResponse(ctx, createdTodo)
}

func (h *TodoHttpHandler) getTodos(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	todos, err := h.todoService.todoRepository.GetTodos(user.Id)
	if err != nil {
		return utility.ErrorResponse(ctx, err)
	}
	return utility.OkResponse(ctx, todos)
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