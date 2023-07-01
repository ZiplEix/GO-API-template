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

// @Summary Get one todo.
// @Description fetch one todo by id.
// @Tags todos
// @Accept */*
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} TodoDB
// @Router /todos/{id} [get]
func (t *TodoController) GetOne(c *fiber.Ctx) error {
	// get the todo id
	id := c.Params("id")

	// get the todo
	todo, err := t.storage.GetTodoByID(id, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail fetching todo",
		})
	}

	// return the todo
	return c.Status(fiber.StatusOK).JSON(todo)
}

// @Summary Update one todo.
// @Description updates one todo by id.
// @Tags todos
// @Accept */*
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200
// @Router /todos/{id} [put]
func (t *TodoController) Update(c *fiber.Ctx) error {
	// get the todo id
	id := c.Params("id")

	todo, err := t.storage.GetTodoByID(id, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail fetching todo",
		})
	}

	// parse the request body
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// update the todo
	todo, err = t.storage.UpdateTodo(*todo, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail updating todo",
		})
	}

	// return the todo
	return c.Status(fiber.StatusOK).JSON(todo)
}

// @Summary Delete one todo.
// @Description deletes one todo by id.
// @Tags todos
// @Accept */*
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200
// @Router /todos/{id} [delete]
func (t *TodoController) Delete(c *fiber.Ctx) error {
	// get the todo id
	id := c.Params("id")

	todo, err := t.storage.GetTodoByID(id, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail fetching todo",
		})
	}

	// delete the todo
	err = t.storage.DeleteTodo(*todo, c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Fail deleting todo",
		})
	}

	// return the todo
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted",
	})
}
