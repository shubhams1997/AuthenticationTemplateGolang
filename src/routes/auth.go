package routes

import (
	"server/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
	// auth.Delete("/delete/:id", controllers.IsAuthorized, controllers.DeleteUser)
	auth.Get("/signout", controllers.Signout)
}
