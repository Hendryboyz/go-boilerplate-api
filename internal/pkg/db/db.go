package db

import (
	"fmt"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {
	connectionString := viper.GetString("db.dsn")
	db, err := gorm.Open(postgres.Open(connectionString))

	if err != nil {
		return &Database{}, fmt.Errorf("db init failed: %w", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(
			"failed to connect database",
			log.String("reason", err.Error()),
		)
	}
	maxConnection := viper.GetInt("db.maxConnection")
	sqlDb.SetMaxOpenConns(maxConnection)

	if viper.GetBool("db.autoMigration") {
		log.Warn("auto migrate db")
		AutoMigrate(db)
	}

	return &Database{
		Client: db,
	}, nil
}
