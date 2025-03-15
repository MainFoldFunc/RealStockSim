package handlers

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/structs"
	"github.com/gofiber/fiber/v2"
)

func CreateStockHandler(c *fiber.Ctx) error {
	log.Println("CreateStockHandler called")

	// Check session cookie
	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No session cookie found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No cookie"})
	}

	// Extract user ID from session
	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil {
		log.Println("Invalid session cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid cookie"})
	}

	// Parse request body into stock struct
	var stock structs.Stocks
	if err := c.BodyParser(&stock); err != nil {
		log.Println("Invalid body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	// Ensure stock is associated with the user
	stock.ID = userID

	// Save stock to database
	if err := database.CreateStockDatabase(&stock); err != nil {
		log.Println("Error in CreateStockDatabase: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	// Update user's portfolio with the created stock
	if err := database.UpdatePortfolioWithStock(userID, stock.Name, stock.AllAmount); err != nil {
		log.Println("Error updating portfolio: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update portfolio"})
	}

	return c.JSON(fiber.Map{
		"message": "Stock created successfully and added to portfolio",
	})
}
