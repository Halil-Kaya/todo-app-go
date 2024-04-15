package user

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"todo/app/validation"
)

type UserHttpHandler struct {
	userService UserService
	logger      *zap.SugaredLogger
}

func NewUserHttpHandler(userService UserService, logger *zap.SugaredLogger) *UserHttpHandler {
	return &UserHttpHandler{userService, logger}
}

func (handler *UserHttpHandler) createUser(ctx *fiber.Ctx) error {
	var request UserCreateDto
	err := ctx.BodyParser(&request)
	if err != nil {
		handler.logger.Fatal(err)
		return err
	}

	if errors := validation.Validate(request); len(errors) > 0 {
		return ctx.JSON(errors)
	}

	createdUser, err := handler.userService.CreateUser(request)
	if err != nil {
		handler.logger.Fatal(err)
		return err
	}
	fmt.Println(createdUser)
	return ctx.JSON(struct {
		Id       string `json:"id"`
		Nickname string `json:"nıckname"`
	}{
		Id:       createdUser.Id.String(),
		Nickname: createdUser.Nickname,
	})
}

func (h *UserHttpHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/user")
	appGroup.Post("/create", h.createUser)
}
