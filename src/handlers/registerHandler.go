package handlers

import (
	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/structs"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RegisterHandler(c *fiber.Ctx) error {
	log.Println("RegisterHandler called")
	var user structs.Users

	if err := c.BodyParser(&user); err != nil {
		log.Println("Invalid reqest body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid reqest body"})
	}

	if err := database.RegisterDatabase(&user); err != nil {
		log.Println("Invalid reqest body")
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Error while adding user to database"})
	}

	return c.JSON(fiber.Map{
		"message": "User succesfully registered",
	})
}
