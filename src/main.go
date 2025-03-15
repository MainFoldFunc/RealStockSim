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
	server.Post("/users/removeUser", handlers.RemoveUserHandler)

	server.Post("/portfolio/createPortfolio", handlers.CreatePortfolioHandler)
	server.Post("/portfolio/deletePortfolio", handlers.DeletePortfolioHandler)

	server.Post("/stock/createStock", handlers.CreateStockHandler)

	fmt.Println("Server starting at port 8080")
	server.Listen(":8080")
}
