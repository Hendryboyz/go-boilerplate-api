/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"go-boilerplate-api/api"
	v1 "go-boilerplate-api/internal/app/api/v1"
	"go-boilerplate-api/internal/pkg/db"
	"go-boilerplate-api/internal/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
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
		log.Fatal("fail to connect database")
	}

	ginRouter := gin.Default()
	defer func() {
		sqlDB, _ := db.Client.DB()
		if closingErr := sqlDB.Close(); closingErr != nil {
			log.Fatal(err.Error())
		} else {
			log.Info("db connection closed")
		}
		log.Sync()
	}()

	ginRouter.Use(ginzap.RecoveryWithZap(log.Default(), true))

	setSwagger()
	setCORS(ginRouter)

	apiV1Router := ginRouter.Group("/v1")
	v1.RegisterRouterApiV1(apiV1Router, db)

	ginRouter.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginRouter.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	port := viper.GetInt32("server.port")
	startEndpoint := fmt.Sprintf("localhost:%d", port)
	server := &http.Server{
		Addr:    startEndpoint,
		Handler: ginRouter.Handler(),
	}

	log.Info(fmt.Sprintf("server starting listen at %s", startEndpoint))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("server failed to listen at %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Warn("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("server forced to shutdown: %s\n", err))
	}
	log.Warn("server exiting")
}

func setSwagger() {
	api.SwaggerInfo.Title = "Todo List API API"
	api.SwaggerInfo.Description = "This API server for a todo app."
	api.SwaggerInfo.Version = "1.0"
	api.SwaggerInfo.BasePath = "/v1"
	api.SwaggerInfo.Schemes = []string{"http", "https"}
}

func setCORS(server *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost",
		"http://google.com",
	}
	// config.AllowAllOrigins = true
	server.Use(cors.New(config))
}
