package middlewares

import (
	"net/http"
	"os"
	"server/src/models"
	"server/src/storage"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(c *fiber.Ctx) error {

	cookie := c.Cookies("jwtk")
	if cookie == "" {
		c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Unauthorized", "success": false, "data": nil})
		return nil
	}
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Unauthorized", "success": false, "data": nil})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expired token
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Unauthorized", "success": false, "data": nil})
		}

		var user models.User
		result := storage.DB.Where("id = ?", claims["id"]).First(&user)
		if result.Error != nil {
			return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Unauthorized", "success": false, "data": nil})
		}
		// c.Set("user", user)

		return c.Next()
	}

	return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{"message": "Unauthorized", "success": false, "data": nil})
}
