package db

import (
	"go-boilerplate-api/internal/model"
	"go-boilerplate-api/internal/pkg/log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal(
			"fail to migrate db",
			log.String("reason", err.Error()),
		)
	}
}
