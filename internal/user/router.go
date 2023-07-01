package user

import (
	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(app *fiber.App, controller *UserController) {

	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Get("/private", controller.AuthentificateUser, Private)
}

func Private(c *fiber.Ctx) error {
	user := c.Locals("user").(*UserDB)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": user,
	})
}
