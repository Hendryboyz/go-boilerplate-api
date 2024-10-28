/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	internalDb "go-boilerplate-api/internal/pkg/db"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "DB migration command",
	Long: `Use this command to migrate all ORM model definitions to database.
For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: startDbMigration,
}

func init() {
	migrationCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startDbMigration(cmd *cobra.Command, args []string) {
	connectionString := viper.GetString("db.dsn")
	db, err := gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Error("db init failed", log.String("reason", err.Error()))
	}

	log.Warn("auto migrate db")
	internalDb.AutoMigrate(db)

	log.Info("db migration success")
}
