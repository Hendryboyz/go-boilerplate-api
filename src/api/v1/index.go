package v1

import "github.com/gin-gonic/gin"

func RegisterRouterApiV1(router *gin.RouterGroup) {
	recordsGroup := router.Group("/todos")
	SetTodoRoutes(recordsGroup)
}
