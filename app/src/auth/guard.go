package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
	"todo/app/core/exception"
	"todo/app/core/utility"
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
			return utility.ErrorResponse(ctx, exception.NewUnauthorized())
		}
		token := strings.Replace(bearToken[0], "Bearer ", "", -1)
		if token == "" {
			return utility.ErrorResponse(ctx, exception.NewUnauthorized())
		}
		userId, err := authGuard.authService.ValidateToken(token)
		if err != nil {
			return utility.ErrorResponse(ctx, exception.NewUnauthorized())
		}
		user := authGuard.authService.userService.FindById(userId)
		if user == nil {
			return utility.ErrorResponse(ctx, exception.NewUnauthorized())
		}
		ctx.Locals("user", user)
		return handler(ctx)
	}
}
