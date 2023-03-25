package controllers

import (
	"net/http"
	"server/src/models"
	"server/src/storage"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	result := storage.DB.Find(&users)
	if result.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Error getting users", "success": false, "data": nil})
		return result.Error
	}
	return c.JSON(users)
}
