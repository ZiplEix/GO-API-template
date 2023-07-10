package router

import (
	"github.com/ZiplEix/API_template/handlers"
	"github.com/ZiplEix/API_template/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	Ping(app)

	Todo(app)

	app.Post("/register", handlers.Signup)
	app.Post("/login", handlers.Login)
	app.Get("/private", middleware.AuthentificateUser, handlers.Private)
}
