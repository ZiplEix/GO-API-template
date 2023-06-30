package router

import "github.com/gofiber/fiber/v2"

func Ping(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
