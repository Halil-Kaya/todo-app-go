package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"todo/app/user"
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
	userHttpHandler := user.NewUserHttpHandler(*userService, logger)

	app := fiber.New()
	app.Use(recover2.New())

	userHttpHandler.RegisterRoutes(app)

	app.Use(cors.New(cors.ConfigDefault))

	port := fmt.Sprintf(":%s", appConfig.Server.Port)
	if err := app.Listen(port); err != nil {
		fmt.Println("There is an error. Server is not up port: ", appConfig.Server.Port, " err : ", err)
	}
}
