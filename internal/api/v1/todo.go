package v1

import (
	"go-boilerplate-api/internal/todo"

	"github.com/gin-gonic/gin"
)

func SetTodoRoutes(router *gin.RouterGroup) {
	controller := todo.ConstructController()

	router.POST("", controller.Create)
	router.GET("", controller.FindAll)
	router.GET(":id", controller.GetItem)
	router.DELETE(":id", controller.Delete)
}
