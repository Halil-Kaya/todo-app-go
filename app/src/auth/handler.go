package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"todo/app/core/utility"
	"todo/app/core/validation"
)

type AuthHttpHandler struct {
	authService AuthService
	logger      *zap.SugaredLogger
}

func NewAuthHttpHandler(authService AuthService, logger *zap.SugaredLogger) *AuthHttpHandler {
	return &AuthHttpHandler{authService, logger}
}

func (h *AuthHttpHandler) login(ctx *fiber.Ctx) error {
	var request LoginDto
	err := ctx.BodyParser(&request)
	if err != nil {
		h.logger.Error(err)
		return utility.ErrorResponse(ctx, err)
	}

	if errors := validation.Validate(request); len(errors) > 0 {
		return utility.ErrorResponse(ctx, errors)
	}
	loginAck, err := h.authService.Login(request)
	if err != nil {
		h.logger.Error(err)
		return utility.ErrorResponse(ctx, err)
	}

	return utility.OkResponse(ctx, loginAck)
}

func (h *AuthHttpHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/auth")
	appGroup.Post("/login", h.login)
}
