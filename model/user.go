package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255); not null; unique"`
	Password string `json:"password" gorm:"type:varchar(255); not null"`
	Todos    []Todo
}
