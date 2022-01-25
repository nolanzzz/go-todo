package model

import (
	"github.com/jinzhu/gorm"
)

type TodoModel struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Completed   int    `json:"completed" gorm:"default:0"`
	TimeSpent   int    `json:"time_spent" gorm:"comment:'Total minutes spent'"`
}

type TodoResponseModel struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   int    `json:"completed"`
	TimeSpent   int    `json:"time_spent"`
}
