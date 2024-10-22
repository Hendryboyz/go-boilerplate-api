package db

import (
	"fmt"
	"log"

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
		log.Fatal("Failed to connect database: ", err)
	}
	maxConnection := viper.GetInt("db.maxConnection")
	sqlDb.SetMaxOpenConns(maxConnection)

	if enableAutoMigration := viper.GetBool("db.autoMigration"); enableAutoMigration == true {
		fmt.Println("auto migrate db")
	}

	return &Database{
		Client: db,
	}, nil
}
