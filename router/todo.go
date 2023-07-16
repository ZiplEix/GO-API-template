package router

import (
	"github.com/ZiplEix/API_template/handlers"
	"github.com/ZiplEix/API_template/middleware"
	"github.com/gofiber/fiber/v2"
)

func Todo(app *fiber.App) {
	todos := app.Group("/todos")

	todos.Use(middleware.AuthentificateUser)

	todos.Get("/", handlers.TodoGetAll)
	todos.Get("/:id", handlers.TodoGetOne)
	todos.Post("/", handlers.TodoCreate)
	todos.Put("/:id", handlers.TodoUpdate)
	todos.Delete("/:id", handlers.TodoDelete)
}
