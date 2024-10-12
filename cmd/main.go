package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hansini0813/divrhino-trivia/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setuoRoutes(app)

	app.Listen(":3000")
}
