package v1

import (
	"go-boilerplate-api/internal/pkg/db"
	"go-boilerplate-api/internal/pkg/log"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func RegisterRouterApiV1(router *gin.RouterGroup, db *db.Database) {
	recordsGroup := router.Group("/todos")
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	recordsGroup.Use(ginzap.Ginzap(log.Default(), time.RFC3339, false))
	// recordsGroup.Use(middleware.SetRequestLogger())
	SetTodoRoutes(recordsGroup, db)
}
