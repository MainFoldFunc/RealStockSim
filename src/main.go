package main

import (
	"fmt"
	"github.com/MainFoldFunc/RealStockSim/src/database"
	"github.com/MainFoldFunc/RealStockSim/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Init()
	server := fiber.New()

	server.Get("/helloworld", handlers.Helloworldhandler)
	server.Post("/users/registerUser", handlers.RegisterHandler)
	server.Post("/users/loginUser", handlers.LoginHandler)
	server.Post("/users/logoutUser", handlers.LogoutHandler)

	server.Post("/portfolio/createPortfolio", handlers.CreatePortfolioHandler)

	fmt.Println("Server starting at port 8080")
	server.Listen(":8080")
}
