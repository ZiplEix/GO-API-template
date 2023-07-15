package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ZiplEix/API_template/database"
	"github.com/ZiplEix/API_template/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthentificateUser(c *fiber.Ctx) error {
	tokenString := c.Cookies("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// decode/validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Status(http.StatusUnauthorized).SendString("Unauthorized")
			return err
		}

		// find user
		var user models.User
		database.DB.Db.Where("id = ?", claims["sub"]).First(&user)
		if user.ID == 0 {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
			return err
		}

		// attach to request
		c.Locals("user", user)

		// continue
		return c.Next()
	} else {
		c.Status(http.StatusUnauthorized).SendString("Unauthorized")
		return err
	}
}
