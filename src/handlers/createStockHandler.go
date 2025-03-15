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

	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No session cookie found")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "No cookie"})
	}

	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil {
		log.Println("Invalid session cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid cookie"})
	}

	var stock structs.Stocks
	if err := c.BodyParser(&stock); err != nil {
		log.Println("Invalid body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	stock.ID = userID

	if err := database.CreateStockDatabase(&stock); err != nil {
		log.Println("Error in CreateStockDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal server error"})
	}

	return c.JSON(fiber.Map{
		"message": "created stock succesfully",
	})

}
