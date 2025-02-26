package todoController

import (
	"gofirst/database"
	"gofirst/repositories"

	"github.com/gofiber/fiber/v2"
)

func ToggleTodoStatus(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := repositories.GetTodoById(id)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	todo.Completed = !todo.Completed
	database.DB.Save(&todo)
	return context.JSON(todo)
}
