package handlers

import (
	"github.com/ZiplEix/API_template/database"
	"github.com/ZiplEix/API_template/models"
	"github.com/gofiber/fiber/v2"
)

// @summary Get all todos
// @description fetch every todo available.
// @tags todos
// @accept */*
// @produce application/json
// @success 200 {object} []models.Todo
// @router /todos [get]
func TodoGetAll(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	todos := []models.Todo{}

	database.DB.Db.Find(&todos, "user_id = ?", userId)

	return c.Status(fiber.StatusOK).JSON(todos)
}

// @summary Get a sigle todo
// @description fetch a single todo by id.
// @tags todos
// @Param id path int true "Todo ID"
// @Produce application/json
// @Success 200 {object} models.Todo
// @Router /todos/:id [get]
func TodoGetOne(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	id := c.Params("id")

	todo := models.Todo{}

	database.DB.Db.Where("user_id = ? AND id = ?", userId, id).First(&todo)

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

type TodoCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// @Summary Create a todo.
// @Description create a single todo.
// @Tags todos
// @Accept json
// @Param todo body TodoCreateRequest true "Todo to create"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos [post]
func TodoCreate(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	var req TodoCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// create the todo
	todo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userId,
	}

	database.DB.Db.Create(&todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}

// @Summary Update a todo.
// @Description update a single todo.
// @Tags todos
// @Accept json
// @Param todo body models.Todo true "Todo update data"
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos/:id [put]
func TodoUpdate(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	id := c.Params("id")

	todo := models.Todo{}

	database.DB.Db.Where("user_id = ? AND id = ?", userId, id).First(&todo)

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	database.DB.Db.Save(&todo)

	return c.Status(fiber.StatusOK).JSON(todo)
}

// @Summary Delete a single todo.
// @Description delete a single todo by id.
// @Tags todos
// @Param id path string true "Todo ID"
// @Produce json
// @Success 200 {object} models.Todo
// @Router /todos/:id [delete]
func TodoDelete(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	id := c.Params("id")

	todo := models.Todo{}

	database.DB.Db.Where("user_id = ? AND id = ?", userId, id).First(&todo)

	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	database.DB.Db.Delete(&todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted",
	})
}
