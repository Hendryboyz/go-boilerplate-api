package db

import (
	"go-boilerplate-api/internal/app/todo/entities"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&entities.Todo{}); err != nil {
	}
}
