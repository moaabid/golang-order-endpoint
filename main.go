package main

import (
	"githu.com/moaabid/golang-order-endpoint/database"
	"githu.com/moaabid/golang-order-endpoint/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	setupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {

	app.Post("/api/user", routes.CreateUser)
}
