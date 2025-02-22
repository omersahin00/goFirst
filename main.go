package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Complated bool   `json:"complated"`
}

var todos = []todo{
	{ID: "1", Item: "Clean room", Complated: false},
	{ID: "2", Item: "Read book", Complated: false},
	{ID: "3", Item: "Record video", Complated: false},
}

func getTodos(context *fiber.Ctx) error {
	return context.JSON(todos)
}

func addTodo(context *fiber.Ctx) error {
	var newTodo todo
	if err := context.BodyParser(&newTodo); err != nil {
		return err
	}

	todos = append(todos, newTodo)
	return context.Status(fiber.StatusCreated).JSON(newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := getTodoById(id)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return context.JSON(todo)
}

func toggleTodoStatus(context *fiber.Ctx) error {
	id := context.Params("id")
	todo, err := getTodoById(id)

	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	todo.Complated = !todo.Complated
	return context.JSON(todo)
}

func main() {
	app := fiber.New()

	app.Get("/todos", getTodos)
	app.Get("/todos/:id", getTodo)
	app.Post("/todos", addTodo)
	app.Patch("/todos/:id", toggleTodoStatus)

	app.Listen(":4000")
}
