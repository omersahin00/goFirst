package controllers

import (

	// Aşağıdaki gibi importlara isim atayabiliyoruz:
	database "gofirst/database"
	models "gofirst/models"
	repositories "gofirst/repositories"

	validator "github.com/go-playground/validator"
	fiber "github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetTodos(context *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return context.JSON(todos)
}

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

func GetTodo(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := repositories.GetTodoById(id)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}
	return context.JSON(todo)
}

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
