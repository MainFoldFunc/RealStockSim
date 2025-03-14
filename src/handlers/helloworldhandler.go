package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Helloworldhandler(c *fiber.Ctx) error {
	return c.SendString("Hello from helloworld handler")
}
