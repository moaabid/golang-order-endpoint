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
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("simple gofiber endpoint")
	})

	///Users Endpoint
	app.Post("/api/createuser", routes.CreateUser)

	app.Get("/api/users", routes.GetUsers)

	app.Get("/api/user/:id", routes.GetUser)

	app.Post("/api/updateuser/:id", routes.UpdateUser)

	app.Delete("/api/deleteuser/:id", routes.DeleteUser)
}
