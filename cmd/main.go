// This declares the package that this file belongs to. In Go, every Go file must belong to a package.
// The main package is special in Go because it tells the Go compiler that this is the entry point of the program, meaning this is where execution starts.
// the main() function inside this pakcage will be he first function to run when the program is executed.
package main

import (
	// imports the fiber package . popular web framework!  It’s used to handle web requests and responses
	"github.com/gofiber/fiber/v2"
	// responsible for handling the database connection logic.
	"github.com/gofiber/template/html/v2"
	"github.com/hansini0813/trivia-api/database"
	"github.com/hansini0813/trivia-api/handlers"
)

// main funciton is required to be inside the main package.
func main() {
	// This calls a function named ConnectDb() from the database package that you imported earlier.
	database.ConnectDb()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	// Set up 404 page
	app.Use(handlers.NotFound)

	app.Listen(":3000")
}
