package todo

import "github.com/gofiber/fiber/v2"

func AddTodoRoutes(app *fiber.App, controller *TodoController) {
	todos := app.Group("/todos")

	todos.Post("/", controller.Create)
	todos.Get("/", controller.GetAll)
	todos.Get("/:id", controller.GetOne)
	todos.Put("/:id", controller.Update)
	todos.Delete("/:id", controller.Delete)
}
