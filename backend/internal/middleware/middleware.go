package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		cost := time.Since(start)
		status := c.Writer.Status()

		// 只记录关键信息，过滤掉健康检查等不重要的请求
		if path != "/health" {
			if status >= 400 {
				// 错误请求详细记录
				logger.Error(fmt.Sprintf("Request error, status: %d, method: %s, path: %s, IP: %s, duration: %v",
					status, c.Request.Method, path, c.ClientIP(), cost))
			} else {
				// 正常请求简洁记录
				logger.Info(fmt.Sprintf("Request, status: %d, method: %s, path: %s, duration: %v",
					status, c.Request.Method, path, cost))
			}
		}
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
