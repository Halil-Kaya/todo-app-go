package user

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"todo/app/core/utility"
	"todo/app/core/validation"
	"todo/app/src/auth"
	"todo/app/src/model"
)

type UserHttpHandler struct {
	userService UserService
	authGuard   auth.AuthGuard
	logger      *zap.SugaredLogger
}

func NewUserHttpHandler(userService UserService, authGuard auth.AuthGuard, logger *zap.SugaredLogger) *UserHttpHandler {
	return &UserHttpHandler{userService, authGuard, logger}
}

func (h *UserHttpHandler) createUser(ctx *fiber.Ctx) error {
	var request UserCreateDto
	err := ctx.BodyParser(&request)
	if err != nil {
		h.logger.Error(err)
		return utility.ErrorResponse(ctx, err)
	}

	if errors := validation.Validate(request); len(errors) > 0 {
		return utility.ErrorResponse(ctx, errors)

	}

	createdUser, err := h.userService.CreateUser(request)
	if err != nil {
		h.logger.Error(err)
		return utility.ErrorResponse(ctx, err)
	}

	ack := UserCreateAck{
		Id:       createdUser.Id.Hex(),
		Nickname: createdUser.Nickname,
	}

	return utility.OkResponse(ctx, ack)
}

func (h *UserHttpHandler) me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	ack := UserMeAck{
		Id:       user.Id.Hex(),
		Nickname: user.Nickname,
		FullName: user.FullName,
	}
	return utility.OkResponse(ctx, ack)
}

func (h *UserHttpHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/user")
	appGroup.Post("/create", h.createUser)
	appGroup.Get("/me", h.authGuard.JWTGuard(h.me))
}
