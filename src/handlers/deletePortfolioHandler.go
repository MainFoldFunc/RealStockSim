package handlers

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/gofiber/fiber/v2"
)

func DeletePortfolioHandler(c *fiber.Ctx) error {
	log.Println("DeletePortfolioHandler called")

	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No session cookie found")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "No cookie found"})
	}

	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil || userID == 0 {
		log.Println("Invalid session cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid cookie"})
	}

	if err := database.DeletePortfolioDatabase(userID); err != nil {
		log.Println("error in DeletePortfolioDatabase")
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Internal server error"})
	}

	return c.JSON(fiber.Map{
		"message": "Portfolio deleted succesfully",
	})

}
