package routes

import (
	"server/src/controllers"
	"server/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterGamesRoutes(app *fiber.App) {
	auth := app.Group("/api/games")
	auth.Post("/", middlewares.IsAuthorized, controllers.GetGames)
}
