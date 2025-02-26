package main

import (
	todoController "gofirst/controllers/todo"
	"gofirst/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.InitDatabase()

	app.Get("/todos", todoController.GetTodos)
	app.Get("/todos/:id", todoController.GetTodo)
	app.Post("/todos", todoController.AddTodo)
	app.Patch("/todos/:id", todoController.ToggleTodoStatus)

	app.Listen(":4000")
}
