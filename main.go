package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	middlewareLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"todo/app/core/middleware"
	"todo/app/src/auth"
	"todo/app/src/todo"
	"todo/app/src/user"
	"todo/config"
	"todo/logger"
	"todo/mongo"
)

func main() {
	appConfig := config.New()
	mongoDb := mongo.ConnectDb(appConfig)
	logger := logger.GetLogger()
	userRepository, err := user.NewUserRepository(mongoDb, logger)
	if err != nil {
		logger.Fatal(err)
	}
	userService := user.NewUserService(*userRepository, logger)

	authService := auth.NewAuthService(userService, logger, appConfig)
	authGuard := auth.NewAuthGuard(*authService, logger)
	authHttpHandler := auth.NewAuthHttpHandler(*authService, logger)

	userHttpHandler := user.NewUserHttpHandler(userService, *authGuard, logger)
	todoHttpHandler := todo.NewTodoHttpHandler(*authGuard)

	app := fiber.New()
	app.Use(middleware.ReqId)
	app.Use(middlewareLogger.New(middlewareLogger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))
	app.Use(recover.New())
	app.Use(cors.New(cors.ConfigDefault))

	userHttpHandler.RegisterRoutes(app)
	authHttpHandler.RegisterRoutes(app)
	todoHttpHandler.RegisterRoutes(app)

	port := fmt.Sprintf(":%s", appConfig.Server.Port)
	if err := app.Listen(port); err != nil {
		fmt.Println("There is an error. Server is not up port: ", appConfig.Server.Port, " err : ", err)
	}
}
