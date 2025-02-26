package todoController

import (
	"gofirst/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetTodo(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := repositories.GetTodoById(id)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}
	return context.JSON(todo)
}
