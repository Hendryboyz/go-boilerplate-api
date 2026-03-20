package v1

import (
	"go-boilerplate-api/internal/app/api/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRouterApi(
	router *gin.RouterGroup,
	todoController *handlers.TodoController,
) {
	todoGroup := router.Group("/todo")
	setTodoRoutes(todoGroup, todoController)
}
