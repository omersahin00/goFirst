package repositories

import (
	"errors"
	"gofirst/database"
	"gofirst/models"
)

func GetTodoById(id string) (*models.Todo, error) {
	var todo models.Todo
	result := database.DB.First(&todo, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("todo not found")
	}
	return &todo, nil
}
