package routes

import (
	"gorm-relation-tutorial/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {

	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.CreateUser)

	app.Get("/lockers", controllers.LockerGetAll)
	app.Post("/lockers", controllers.CreateLocker)

}
