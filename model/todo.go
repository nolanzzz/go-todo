package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(255); not null"`
	Description string `json:"description" gorm:"type:text; null"`
	Completed   int    `json:"completed" gorm:"default:0"`
	TimeSpent   int    `json:"time_spent" gorm:"comment:'Total minutes spent'"`
	UserID      uint   `json:"user_id"`
	User        User
}

type TodoResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   int    `json:"completed"`
	TimeSpent   int    `json:"time_spent"`
}
