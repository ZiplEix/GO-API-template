package todo

import "github.com/gofiber/fiber/v2"

// Usage:
// Ce package fournit un contrôleur pour gérer les requêtes HTTP liées aux tâches à faire ("todos").
// Il définit les structures de données pour les requêtes et les réponses, et les méthodes de gestion des requêtes pour créer des todos.
//
// Pour utiliser ce contrôleur, instanciez-le avec un `TodoStorage`, puis enregistrez ses méthodes comme gestionnaires de routes avec votre serveur HTTP.
// Le contrôleur utilise le framework Fiber (https://github.com/gofiber/fiber) pour la gestion des requêtes.
//
// Exemple:
//    app := fiber.New()
//    todoController := NewTodoController(todoStorage)
//    app.Post("/todos", todoController.Create)
//    app.Listen(":3000")
//
// Note: Les annotations @Summary, @Description, etc., sont utilisées pour la génération automatique de la documentation de l'API avec Swagger (https://swagger.io/).

type TodoController struct {
	storage *TodoStorage
}

func NewTodoController(storage *TodoStorage) *TodoController {
	return &TodoController{storage: storage}
}

type createTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type createTodoResponse struct {
	ID string `json:"id"`
}

// @Summary Create one todo.
// @Description creates one todo.
// @Tags todos
// @Accept */*
// @Produce json
// @Param todo body createTodoRequest true "Todo to create"
// @Success 200 {object} createTodoResponse
// @Router /todos [post]
func (t *TodoController) Create(c *fiber.Ctx) error {
	// parse the request body
	var req createTodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// create the todo
	id, err := t.storage.CreateTodo(
		req.Title,
		req.Description,
		false,
		c.Context(),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail creating todo",
		})
	}

	// return the todo id
	return c.Status(fiber.StatusOK).JSON(createTodoResponse{
		ID: id,
	})
}

// @Summary Get all todos.
// @Description fetch every todo available.
// @Tags todos
// @Accept */*
// @Produce json
// @Success 200 {object} []TodoDB
// @Router /todos [get]
func (t *TodoController) GetAll(c *fiber.Ctx) error {
	// get all todos
	todos, err := t.storage.GetAllTodos(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail fetching todos",
		})
	}

	// return the todos
	return c.Status(fiber.StatusOK).JSON(todos)
}
