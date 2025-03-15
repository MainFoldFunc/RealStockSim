package handlers

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/structs"
	"github.com/gofiber/fiber/v2"
)

func CreatePortfolioHandler(c *fiber.Ctx) error {
	log.Println("CreatePortfolioHandler called")

	// Retrieve the session cookie
	sessionCookie := c.Cookies("session")
	if sessionCookie == "" {
		log.Println("No session cookie found")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "User not logged in"})
	}

	// Parse the session cookie value as the user ID
	var userID uint
	_, err := fmt.Sscanf(sessionCookie, "%d", &userID)
	if err != nil || userID == 0 {
		log.Println("Invalid session cookie")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "Invalid session"})
	}

	// Now that we have the userID, we can create the portfolio
	var portfolio structs.Portfolio
	if err := c.BodyParser(&portfolio); err != nil {
		log.Println("Invalid body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	// Associate the userID with the portfolio
	portfolio.UserID = userID

	// Save the portfolio to the database
	if err := database.CreatePortfolioDatabase(&portfolio); err != nil {
		log.Println("Error in CreatePortfolioDatabase:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Error while creating portfolio"})
	}

	return c.JSON(fiber.Map{
		"message": "Portfolio successfully created",
	})
}
