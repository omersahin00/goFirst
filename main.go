package main

import (
	"gofirst/controllers"
	"gofirst/database"

	"github.com/gofiber/fiber/v2"
)

// main.go
func main() {
	app := fiber.New()

	database.InitDatabase()

	app.Get("/todos", controllers.GetTodos)
	app.Get("/todos/:id", controllers.GetTodo)
	app.Post("/todos", controllers.AddTodo)
	app.Patch("/todos/:id", controllers.ToggleTodoStatus)

	app.Listen(":4000")
}
