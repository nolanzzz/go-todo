package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null, unique"`
	Password string `json:"password" gorm:"not null"`
	Todos    []Todo
}
