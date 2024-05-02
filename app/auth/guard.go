package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
)

type AuthGuard struct {
	authService AuthService
	logger      *zap.SugaredLogger
}

func NewAuthGuard(authService AuthService, logger *zap.SugaredLogger) *AuthGuard {
	return &AuthGuard{authService, logger}
}

func (authGuard AuthGuard) JWTGuard(handler fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		headers := ctx.GetReqHeaders()
		bearToken, exists := headers["Authorization"]
		if !exists {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		token := strings.Replace(bearToken[0], "Bearer ", "", -1)
		if token == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		userId := authGuard.authService.ValidateToken(token)
		user := authGuard.authService.userService.FindById(userId)
		if user == nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Locals("user", user)
		return handler(ctx)
	}
}
