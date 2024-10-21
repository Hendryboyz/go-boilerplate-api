package v1

import (
	"go-boilerplate-api/db"
	"go-boilerplate-api/internal/todo"

	"github.com/gin-gonic/gin"
)

func SetTodoRoutes(router *gin.RouterGroup, db *db.Database) {
	controller := todo.ConstructController(db)

	router.POST("", controller.Create)
	router.GET("", controller.FindAll)
	router.GET(":id", controller.GetItem)
	router.DELETE(":id", controller.Delete)
}
