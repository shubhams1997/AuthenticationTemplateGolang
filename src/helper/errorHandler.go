package helper

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleFiberError(err error, message string, c *fiber.Ctx, status int) error {
	// TODO: Add logging
	log.Println(err)
	return c.Status(status).JSON(&fiber.Map{"message": message, "success": false, "data": nil})
}
