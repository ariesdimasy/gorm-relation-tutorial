package main

import (
	"gorm-relation-tutorial/config"
	"gorm-relation-tutorial/migrations"
	"gorm-relation-tutorial/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDatabase()
	migrations.Migration()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	routes.RouteInit(app)

	app.Listen(":6780")
}
