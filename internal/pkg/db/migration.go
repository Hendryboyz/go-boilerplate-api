package db

import (
	"go-boilerplate-api/internal/app/todo/entities"
	"go-boilerplate-api/internal/pkg/log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&entities.Todo{}); err != nil {
		log.Fatal(
			"fail to migrate db",
			log.String("reason", err.Error()),
		)
	}
}
