package handlers

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/structs"
	"github.com/gofiber/fiber/v2"
)

func BuyStockHandler(c *fiber.Ctx) error {
	log.Println("BuyStockHandler called")
	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "No cookie"})
	}

	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil {
		log.Println("invalid session cookie: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid cookie"})
	}

	var whatToBuy structs.BuyingStocks
	if err := c.BodyParser(&whatToBuy); err != nil {
		log.Println("Invalid body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	if err := database.BuyStockDatabase(&whatToBuy, userID); err != nil {
		log.Println("Error in BuyStockDatabase:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "bought stock",
	})

}
