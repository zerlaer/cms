package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter 简单的基于内存的速率限制器
type RateLimiter struct {
	rate      int           // 允许的请求数
	per       time.Duration // 时间窗口
	tokens    map[string][]time.Time // IP地址 -> 请求时间列表
	mu        sync.Mutex
}

// NewRateLimiter 创建一个新的速率限制器
func NewRateLimiter(rate int, per time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:   rate,
		per:    per,
		tokens: make(map[string][]time.Time),
	}
}

// Limit 限制请求频率的中间件
func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		
		rl.mu.Lock()
		defer rl.mu.Unlock()
		
		// 清理过期的请求记录
		now := time.Now()
		times := rl.tokens[ip]
		validTimes := make([]time.Time, 0, len(times))
		
		for _, t := range times {
			if now.Sub(t) < rl.per {
				validTimes = append(validTimes, t)
			}
		}
		
		// 检查是否超过限制
		if len(validTimes) >= rl.rate {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}
		
		// 添加当前请求记录
		validTimes = append(validTimes, now)
		rl.tokens[ip] = validTimes
		
		c.Next()
	}
}
