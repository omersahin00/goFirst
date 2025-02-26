package todoController

import (
	"gofirst/database"
	"gofirst/models"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func AddTodo(context *fiber.Ctx) error {
	var newTodo models.Todo
	if err := context.BodyParser(&newTodo); err != nil {
		return err
	}

	if err := validate.Struct(&newTodo); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newTodo.Completed = false

	database.DB.Create(&newTodo)
	return context.Status(fiber.StatusCreated).JSON(newTodo)
}
