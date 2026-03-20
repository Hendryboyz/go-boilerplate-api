package v1

import (
	"go-boilerplate-api/internal/app/api/http/handlers"

	"github.com/gin-gonic/gin"
)

func setTodoRoutes(router *gin.RouterGroup, customerHandler *handlers.TodoController) {
	router.POST("", customerHandler.CreateItem)
	router.GET("", customerHandler.FindAll)
	router.GET(":itemId", customerHandler.GetItem)
	router.PATCH(":itemId", customerHandler.UpdateItem)
	router.DELETE(":itemId", customerHandler.DeleteItem)
}
