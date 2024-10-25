package v1

import (
	"go-boilerplate-api/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func RegisterRouterApiV1(router *gin.RouterGroup, db *db.Database) {
	recordsGroup := router.Group("/todos")
	SetTodoRoutes(recordsGroup, db)
}