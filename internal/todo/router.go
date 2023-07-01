package todo

import "github.com/gofiber/fiber/v2"

func AddTodoRoutes(app *fiber.App, controller *TodoController) {
	todos := app.Group("/todos")

	todos.Post("/", controller.Create)
	todos.Get("/", controller.GetAll)
	// todos.Get("/todos/:id", controller.Get)
	// todos.Put("/todos/:id", controller.Update)
	// todos.Delete("/todos/:id", controller.Delete)
}
