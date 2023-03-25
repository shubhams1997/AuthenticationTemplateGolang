package routes

import (
	"server/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/users", controllers.GetUsers)
}
