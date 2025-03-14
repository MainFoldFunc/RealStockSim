package main

import (
	"fmt"
	"github.com/MainFoldFunc/RealStockSim/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	server.Get("/helloworld", handlers.Helloworldhandler)
	fmt.Println("Server starting at port 8080")
	server.Listen(":8080")
}
