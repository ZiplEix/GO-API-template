package todo

import (
	"github.com/ZiplEix/API_template/internal/user"
	"github.com/gofiber/fiber/v2"
)

func AddTodoRoutes(app *fiber.App, controller *TodoController, auth *user.UserController) {
	todos := app.Group("/todos", auth.AuthentificateUser)

	todos.Post("/", controller.Create)
	todos.Get("/", controller.GetAll)
	todos.Get("/:id", controller.GetOne)
	todos.Put("/:id", controller.Update)
	todos.Delete("/:id", controller.Delete)
}
