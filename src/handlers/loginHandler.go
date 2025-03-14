package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/structs"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var user structs.LoginUser
	log.Println("LoginUser handler called")

	if err := c.BodyParser(&user); err != nil {
		log.Println("Invalid reqest body")
		return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "Invalid body"})
	}

	storedUser, err := database.GetUserByEmail(user.Email)
	if err != nil {
		log.Println("Error while checking if user egzists: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": "Error while searching if the user egzists"})
	}

	if storedUser.Password != user.Password {
		log.Println("Invalid password")
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "User not authorized"})
	}

	cookie := fiber.Cookie{
		Name:     "session",
		Value:    fmt.Sprintf("%d", storedUser.ID),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
	}

	c.Cookie(&cookie)
	log.Println("User logged int")

	return c.JSON(fiber.Map{
		"message": "User succesfully loggedin",
	})
}
