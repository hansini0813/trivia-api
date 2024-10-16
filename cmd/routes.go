package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hansini0813/trivia-api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	// Add new route for new view
	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)

	// Route to delete all facts
	app.Delete("/facts", handlers.DeleteAllFacts)
	// Route to delete selected facts
	app.Post("/facts/delete", handlers.DeleteSelectedFacts)
}
