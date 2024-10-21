package main

import (
	"fmt"
	v1 "go-boilerplate-api/api/v1"
	"go-boilerplate-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	config := config.Init("local")

	defer func() {}()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"": "OK"})
	})

	apiV1Router := server.Group("/v1")
	v1.RegisterRouterApiV1(apiV1Router)

	port := config.GetInt32("server.port")
	server.Run(fmt.Sprintf("localhost:%d", port))
}
