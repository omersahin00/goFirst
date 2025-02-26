package todoController

import (
	"gofirst/database"
	"gofirst/models"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(context *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return context.JSON(todos)
}
