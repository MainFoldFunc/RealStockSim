package handlers

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/structs"
	"github.com/gofiber/fiber/v2"
)

func UpdateStocksHandler(c *fiber.Ctx) error {
	log.Println("UpdateStocksHandler called")
	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "No cookie"})
	}

	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil {
		log.Println("Invalid session cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid cookie"})
	}

	var stock structs.UpdateStock
	if err := c.BodyParser(&stock); err != nil {
		log.Println("Invalid body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	// Pass userID and stock information to the database function
	stock.ID = userID

	if err := database.UpdateStockDatabase(&stock); err != nil {
		log.Println("Error in UpdateStockDatabase: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal server error"})
	}

	return c.JSON(fiber.Map{
		"message": "Stock updated successfully",
	})
}
