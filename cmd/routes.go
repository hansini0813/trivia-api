package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hansini0813/trivia-api/handlers"
)

func setuoRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
