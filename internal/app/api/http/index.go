package http

import (
	"go-boilerplate-api/docs"
	"go-boilerplate-api/global"
	"go-boilerplate-api/internal/app/api/http/handlers"
	"go-boilerplate-api/internal/app/api/http/middleware"
	v1 "go-boilerplate-api/internal/app/api/http/routes/v1"
	"go-boilerplate-api/internal/pkg/log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgin"
)

func NewApiRouter(
	todoController *handlers.TodoController,
) *gin.Engine {
	var ginRouter *gin.Engine
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		ginRouter = gin.New()
	} else {
		ginRouter = gin.Default()
	}
	if global.App.Config.Server.Profiling {
		pprof.Register(ginRouter)
	}
	ginRouter.Use(
		apmgin.Middleware(ginRouter),
		ginzap.RecoveryWithZap(log.Default(), true),
		middleware.SetUnexpectedPanicsHandler(),
	)
	setCORS(ginRouter)
	setSwagger()

	apiV1Router := ginRouter.Group("/v1")
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	// apiV1Router.Use(
	// 	ginzap.Ginzap(log.Default(), time.RFC3339, false))
	apiV1Router.Use(
		ginzap.GinzapWithConfig(log.Default(), &ginzap.Config{
			UTC:        false,
			TimeFormat: time.RFC3339,
			Context: ginzap.Fn(func(c *gin.Context) []log.Field {
				tx := apm.TransactionFromContext(c.Request.Context())
				traceContext := tx.TraceContext()
				traceId := traceContext.Trace.String()
				// log trace and span ID
				return []log.Field{log.String("traceId", traceId)}
			}),
		}))

	v1.RegisterRouterApi(apiV1Router, todoController)

	ginRouter.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	environment := global.App.Config.Server.Environment

	if environment == "local" {
		ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return ginRouter
}

// const swaggerHost = "ecovision-svc.dev.deltaww-energy.com"

func setSwagger() {
	docs.SwaggerInfo.Title = "Electricity Bill Management API"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = swaggerHost
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func setCORS(server *gin.Engine) {
	config := cors.DefaultConfig()
	allowOrigins := strings.Split(global.App.Config.Server.AllowOrigins, ",")
	if len(allowOrigins) == 0 {
		allowOrigins = []string{
			"http://localhost",
			"http://google.com",
		}
	}

	config.AllowOrigins = allowOrigins
	config.AllowMethods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "Signature"}
	config.AllowCredentials = true
	server.Use(cors.New(config))
}
