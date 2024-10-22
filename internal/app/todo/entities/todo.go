package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	UserId      string
	Description string
}
