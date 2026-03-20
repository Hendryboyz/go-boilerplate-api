package middleware

import (
	"go-boilerplate-api/internal/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUnexpectedPanicsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic occurred", log.Any("error", r))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected server error"})
				c.Abort()
			}
		}()

		c.Next()
	}
}
