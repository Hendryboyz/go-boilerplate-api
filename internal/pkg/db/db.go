package db

import (
	"fmt"
	"go-boilerplate-api/global"
	"go-boilerplate-api/internal/pkg/log"
	"time"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)

func NewRDBMS() *gorm.DB {
	connectionString := global.App.Config.DB.DSN
	createBatchSize := global.App.Config.DB.CreateBatchSize
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		CreateBatchSize: createBatchSize,
	})

	if err != nil {
		panic(fmt.Errorf("db init failed: %w", err))
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(
			"failed to connect database",
			log.String("reason", err.Error()),
		)
	}
	maxConnection := global.App.Config.DB.MaxConnection
	sqlDb.SetMaxOpenConns(maxConnection)
	sqlDb.SetMaxIdleConns(maxConnection / 2)
	sqlDb.SetConnMaxIdleTime(15 * time.Minute)
	return db
}
