package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-boilerplate-api/internal/pkg/log"
	"io"

	"github.com/gin-gonic/gin"
)

func SetRequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		fields := []log.Field{}

		if ctx.Request.Body != nil {
			var buf bytes.Buffer
			tee := io.TeeReader(ctx.Request.Body, &buf)
			bodyBytes, _ := io.ReadAll(tee)
			body := make(map[string]interface{})
			if err := json.Unmarshal(bodyBytes, &body); err == nil {
				fields = append(fields, log.Any("body", body))
			}
			ctx.Request.Body = io.NopCloser(&buf)
		}

		log.Info(
			fmt.Sprintf("[%s] %s", method, path),
			fields...,
		)

		ctx.Next()
	}
}
