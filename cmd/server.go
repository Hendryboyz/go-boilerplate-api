/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-boilerplate-api/docs"
	v1 "go-boilerplate-api/internal/app/api/v1"
	"go-boilerplate-api/internal/pkg/db"
	"go-boilerplate-api/internal/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: startServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func startServer(cmd *cobra.Command, args []string) {
	// connect database
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Can't connect database")
	}

	server := gin.New()

	defer func() {
		sqlDB, _ := db.Client.DB()
		if closingErr := sqlDB.Close(); closingErr != nil {
			log.Fatal(err.Error())
		} else {
			log.Info("Closed db connection")
		}
		log.Sync()
	}()

	docs.SwaggerInfo.Title = "Todo List API API"
	docs.SwaggerInfo.Description = "This API server for a todo app."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	apiV1Router := server.Group("/v1")
	v1.RegisterRouterApiV1(apiV1Router, db)

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	port := viper.GetInt32("server.port")
	startEndpoint := fmt.Sprintf("localhost:%d", port)
	log.Info(fmt.Sprintf("server start at %s", startEndpoint))
	server.Run(startEndpoint)
}
