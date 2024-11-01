package v1

import (
	"go-boilerplate-api/internal/app/todo"
	"go-boilerplate-api/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func SetTodoRoutes(router *gin.RouterGroup, db *db.Database) {
	controller := todo.ConstructController(db)

	router.POST("", controller.CreateItem)
	router.GET("", controller.FindAll)
	router.GET(":itemId", controller.GetItem)
	router.PATCH(":itemId", controller.UpdateItem)
	router.DELETE(":itemId", controller.DeleteItem)
}
