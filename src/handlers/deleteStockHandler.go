package handlers

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/gofiber/fiber/v2"
)

func DeleteStockHandler(c *fiber.Ctx) error {
	log.Println("DeleteStockHandler called")
	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "No cookie"})
	}

	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil {
		log.Println("Invalid cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "invalid cookie"})
	}

	if err := database.DeleteStockDatabase(userID); err != nil {
		log.Println("Error in DeleteStockDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal server error"})
	}

	return c.JSON(fiber.Map{
		"message": "StockDeleted succesfully",
	})

}
