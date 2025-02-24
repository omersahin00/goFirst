package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID        string `json:"id" gorm:"type:uuid;primary_key;"`
	Item      string `json:"item" validate:"required,min=3,max=100"`
	Completed bool   `json:"completed" gorm:"default:false"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	todo.ID = uuid.New().String()
	return
}
