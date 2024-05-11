package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"todo/app/validation"
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
		return err
	}

	if errors := validation.Validate(request); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	loginAck, err := h.authService.Login(request)
	if err != nil {
		return err
	}

	return ctx.JSON(loginAck)
}

func (h *AuthHttpHandler) RegisterRoutes(app *fiber.App) {
	appGroup := app.Group("/auth")
	appGroup.Post("/login", h.login)
}
