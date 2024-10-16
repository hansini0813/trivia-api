package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hansini0813/trivia-api/database"
	"github.com/hansini0813/trivia-api/models"
)

func DeleteSelectedFacts(c *fiber.Ctx) error {
	// Define a struct to capture the IDs from the request body
	type Request struct {
		IDs []uint `json:"ids"`
	}

	var request Request

	// Parse the JSON body to get the selected fact IDs
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	// Check if no IDs were provided
	if len(request.IDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No facts selected for deletion",
		})
	}

	// Delete facts with the provided IDs
	result := database.DB.Db.Where("id IN ?", request.IDs).Delete(&models.Fact{})

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete selected facts",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Selected facts have been deleted",
	})
}

func DeleteAllFacts(c *fiber.Ctx) error {
	// Delete all facts
	result := database.DB.Db.Exec("DELETE FROM facts")

	// Check for errors
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete facts",
			"error":   result.Error.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "All facts have been deleted",
	})
}
func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Trivia Time",
		"Subtitle": "Facts for funtimes with friends!",
		// displaying all facts on the index page!
		"Facts": facts,
	})
}

// Create new Fact View handler
func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return ConfirmationView(c) // 2. Return confirmation view after you have created a fact!
}

// 1. New Confirmation view
func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
		"Link":     "/", // You can include a link that will go back to the homepage
	})
}
