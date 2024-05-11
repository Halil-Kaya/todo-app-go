package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"todo/app/auth"
	"todo/app/model"
	"todo/app/utility"
	"todo/app/validation"
)

type UserHttpHandler struct {
	userService UserService
	authGuard   auth.AuthGuard
	logger      *zap.SugaredLogger
}

func NewUserHttpHandler(userService UserService, authGuard auth.AuthGuard, logger *zap.SugaredLogger) *UserHttpHandler {
	return &UserHttpHandler{userService, authGuard, logger}
}

func (handler *UserHttpHandler) createUser(ctx *fiber.Ctx) error {
	var request UserCreateDto
	err := ctx.BodyParser(&request)
	if err != nil {
		handler.logger.Error(err)
		return utility.ErrorResponse(ctx, err)
	}

	if errors := validation.Validate(request); len(errors) > 0 {
		return utility.ErrorResponse(ctx, errors)

	}

	createdUser, err := handler.userService.CreateUser(request)
	if err != nil {
		handler.logger.Error(err)
		return utility.ErrorResponse(ctx, err)
	}

	ack := struct {
		Id       string `json:"id"`
		Nickname string `json:"nıckname"`
	}{
		Id:       createdUser.Id.Hex(),
		Nickname: createdUser.Nickname,
	}

	return utility.OkResponse(ctx, ack)
}

func (handler *UserHttpHandler) me(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*model.User)
	fmt.Println("gelen user -< ", user)
	return nil
}

func (h *UserHttpHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/user")
	appGroup.Post("/create", h.createUser)
	appGroup.Get("/me", h.authGuard.JWTGuard(h.me))
}
