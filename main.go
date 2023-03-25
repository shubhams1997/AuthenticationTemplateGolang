package main

import (
	"fmt"
	"log"
	"server/src/models"
	"server/src/routes"
	"server/src/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	HandleError(err, "Error loading .env file")

	db, err := storage.NewConnection()
	HandleError(err, "Error connecting to database")

	err = models.MigrateUsers(db)
	HandleError(err, "Error migrating users")

	fmt.Println("Starting server...")
	app := fiber.New()

	// middleware
	// app.Use(middleware.Logger())

	// Register routes
	routes.RegisterAuthRoutes(app)
	routes.RegisterUserRoutes(app)
	routes.RegisterGamesRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

func HandleError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
