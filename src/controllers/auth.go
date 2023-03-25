package controllers

import (
	"net/http"
	"os"
	"server/src/helper"
	"server/src/models"
	"server/src/storage"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	type request struct {
		PhoneNumber string `json:"phoneNumber"`
		Password    string `json:"password"`
	}
	var body request
	err := c.BodyParser(&body)
	if err != nil {
		return helper.HandleFiberError(err, "Error parsing request body", c, http.StatusBadRequest)
	}
	var user models.User
	result := storage.DB.Where("phone_number = ?", body.PhoneNumber).First(&user)

	user.LastLogin = time.Now().UTC()
	user.IsActive = true
	storage.DB.Save(&user)

	if result.Error != nil {
		return helper.HandleFiberError(result.Error, "Invalid Credentials", c, http.StatusNotFound)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["phoneNumber"] = user.PhoneNumber
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return helper.HandleFiberError(err, "Error signing token", c, http.StatusInternalServerError)
	}

	//set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwtk",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"token": t,
		"user": models.User{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			Coins:       user.Coins,
		},
		"message": "Login successful",
		"success": true,
	})

}

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.IsActive = true
	result := storage.DB.Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate") {
			return helper.HandleFiberError(result.Error, "User with this Phone number already exists", c, http.StatusBadRequest)
		}
		return helper.HandleFiberError(result.Error, "Error creating user", c, http.StatusBadRequest)
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	result := storage.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Error deleting user", "success": false, "data": nil})
		return result.Error
	}
	return c.JSON(&fiber.Map{"message": "User deleted successfully", "success": true, "data": nil})
}

func Signout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwtk",
		Expires:  time.Now().Add(-(time.Hour * 1)),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.JSON(&fiber.Map{"message": "Logout successful", "success": true, "data": nil})
}
