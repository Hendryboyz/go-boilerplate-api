/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"go-boilerplate-api/bootstrap"
	"go-boilerplate-api/cmd"
	"go-boilerplate-api/internal/pkg/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dse-electricity-bill-management-service",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.InitConfiguration()
		app, err := wireApp()
		if err != nil {
			panic(err)
		}
		if err := app.Run(); err != nil {
			panic(err)
		}

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Warn("shutting down gracefully, press Ctrl+C again to force")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := app.Shutdown(ctx); err != nil {
			log.Fatal(fmt.Sprintf("server forced to shutdown: %s\n", err))
			panic(err)
		}
		log.Warn("server exiting")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	cmd.RegisterCommands(rootCmd)
	Execute()
}
