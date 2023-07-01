package user

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (u *UserController) AuthentificateUser(c *fiber.Ctx) error {
	tokenString := c.Cookies("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// decode/validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check if the signing method is HMAC
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
			return err
		}

		id := fmt.Sprintf("%v", claims["sub"])

		// find user
		user, err := u.storage.GetUserByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Fail fetching user",
			})
		}

		// set user in locals
		c.Locals("user", user)

		// continue
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
}
