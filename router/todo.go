package router

import (
	"github.com/ZiplEix/API_template/handlers"
	"github.com/gofiber/fiber/v2"
)

func Todo(app *fiber.App) {
	todos := app.Group("/todos")

	todos.Get("/", handlers.TodoGetAll)
	todos.Get("/:id", handlers.TodoGetOne)
	todos.Post("/", handlers.TodoCreate)
	todos.Put("/:id", handlers.TodoUpdate)
	todos.Delete("/:id", handlers.TodoDelete)
}
