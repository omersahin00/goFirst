package controllers

import (
	"errors"

	// Aşağıdaki gibi importlara isim atayabiliyoruz:
	database "gofirst/database"
	models "gofirst/models"

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

// repo olmalı:
func GetTodoById(id string) (*models.Todo, error) {
	var todo models.Todo
	result := database.DB.First(&todo, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}

func GetTodo(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := GetTodoById(id)
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}
	return context.JSON(todo)
}

func ToggleTodoStatus(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := GetTodoById(id)

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
