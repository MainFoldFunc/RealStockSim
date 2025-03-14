package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LogoutHandler(c *fiber.Ctx) error {
	log.Println("LogoutHandler called")
	// Set the session cookie's expiration time to the past to effectively delete it
	cookie := fiber.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now().Add(-24 * time.Hour), // Expire it immediately
		HTTPOnly: true,
		Secure:   false, // You can set this to true if you are using HTTPS
	}

	c.Cookie(&cookie) // Remove the session cookie
	log.Println("User logged out successfully")

	return c.JSON(fiber.Map{
		"message": "User successfully logged out",
	})
}
